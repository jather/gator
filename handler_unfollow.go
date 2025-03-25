package main

import (
	"context"
	"errors"

	"github.com/jather/rss-feed-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("expecting 1 argument")
	}
	url := cmd.arguments[0]
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, Url: url})
	if err != nil {
		return err
	}
	return nil
}
