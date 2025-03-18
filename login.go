package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("expecting one argument")
	}
	username := cmd.arguments[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Username has been set to %s\n", username)
	return nil
}
