package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("expecting one argument")
	}
	username := cmd.arguments[0]

	//check if user exists
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("username hasn't been registered")
		}
	}
	// set username
	err = s.cfg.SetUser(username)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Username has been set to %s\n", username)
	return nil
}
