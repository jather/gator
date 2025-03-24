package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/jather/rss-feed-aggregator/internal/config"
	"github.com/jather/rss-feed-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, _ := config.Read()
	db, err := sql.Open("postgres", cfg.Dburl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	appState := state{cfg: &cfg, db: dbQueries}
	commandList := commands{
		map[string](func(*state, command) error){},
	}
	commandList.register("login", handlerLogin)
	commandList.register("register", handlerRegister)
	commandList.register("reset", handlerReset)
	commandList.register("users", handlerUsers)
	commandList.register("agg", handlerAgg)
	commandList.register("addfeed", handlerAddFeed)
	commandList.register("feeds", handlerFeeds)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough arguments")
	}
	commandName := args[1]
	commandArgs := args[2:]
	cmd := command{commandName, commandArgs}
	err = commandList.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
