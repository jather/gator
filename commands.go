package main

import (
	"errors"
)

type command struct {
	name      string
	arguments []string
}
type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.commands[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	err := function(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
