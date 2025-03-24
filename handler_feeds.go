package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("expecting no arguments")
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(feeds)
	for _, feed := range feeds {
		fmt.Println("\nname:", feed.Name)
		fmt.Println("url:", feed.Url)
		fmt.Println("user:", feed.Username.String)
	}
	return nil
}
