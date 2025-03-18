package main

import (
	"log"
	"os"

	"github.com/jather/rss-feed-aggregator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, _ := config.Read()
	appState := state{&cfg}
	commandList := commands{
		map[string](func(*state, command) error){},
	}
	commandList.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}
	commandName := args[1]
	commandArgs := args[2:]
	cmd := command{commandName, commandArgs}
	err := commandList.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
