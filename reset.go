package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("expecting no arguments")
	}
	err := s.db.ResetUser(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("user table reset in database")
	return nil
}
