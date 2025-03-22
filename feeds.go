package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
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
	req.Header.Set("User-Agent", "rss-feed-aggregator")
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
	for _, item := range feed.Channel.Item {
		unescapeReassign(&item.Title, &item.Description)
	}
	return &feed, nil
}
