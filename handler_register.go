package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jather/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("expecting one argument")
	}
	name := cmd.arguments[0]

	// check if user exists already
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return errors.New("name already exists")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name})
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("User %s was created. %v", user.Name, user)

	return nil
}
