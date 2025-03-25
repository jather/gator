package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/jather/rss-feed-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return errors.New("expecting 0 arguments")
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("feeds for user %s:\n", user.Name)
	for _, feed := range feeds {
		fmt.Println(feed.FeedName.String)
	}
	return nil
}
