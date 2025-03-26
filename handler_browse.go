package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jather/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.arguments) > 1 {
		return errors.New("expecting 1 optional argument")
	} else if len(cmd.arguments) == 1 {
		val, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return errors.New("argument should be integer")
		}
		limit = val
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return err
	}
	fmt.Println("showing posts:")
	for _, post := range posts {
		fmt.Printf("\n%v", post.Title)
		fmt.Printf("\n%v", post.Description.String)
		fmt.Println()
	}

	return nil
}
