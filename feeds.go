package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jather/gator/internal/database"
	"github.com/lib/pq"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func unescapeReassign(fields ...*string) {
	for _, field := range fields {
		*field = html.UnescapeString(*field)
	}
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := http.DefaultClient
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	feed := RSSFeed{}
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}
	unescapeReassign(&feed.Channel.Description, &feed.Channel.Title)
	for i := range feed.Channel.Item {
		unescapeReassign(&feed.Channel.Item[i].Title, &feed.Channel.Item[i].Description)
	}
	return &feed, nil
}

func scrapeFeeds(s *state) error {
	feed_to_fetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{ID: feed_to_fetch.ID, LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}})
	if err != nil {
		return err
	}
	feed, err := fetchFeed(context.Background(), feed_to_fetch.Url)
	if err != nil {
		return err
	}
	fmt.Printf("\nSaving posts from feed: %s\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
		description := sql.NullString{String: item.Description, Valid: true}
		if item.Title == "" {
			continue
		}
		if item.Description == "" {
			description.Valid = false
		}
		publishedAt := sql.NullTime{Time: time.Time{}, Valid: false}
		if pubdate, err := time.Parse(time.RFC3339, item.PubDate); err != nil {
			publishedAt.Time = pubdate
		}
		if item.PubDate == "" {
			description.Valid = false
		}
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Title: item.Title, Url: item.Link, Description: description, PublishedAt: publishedAt, FeedID: feed_to_fetch.ID})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code != "23505" {
					fmt.Println(err)
				}
			}
		}

	}
	return nil
}
