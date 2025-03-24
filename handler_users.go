package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("expecting no arguments")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if s.cfg.CurrentUserName == user {
			fmt.Println("*", user, "(current)")
		} else {
			fmt.Println("*", user)
		}
	}
	return nil
}
