package main

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("expecting 1 argument")
	}
	time_between_reqs := cmd.arguments[0]
	fmt.Println("Collecting feeds every", time_between_reqs)

	interval, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
