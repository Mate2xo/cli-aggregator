package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Mate2xo/gator/internal/config"
	"github.com/Mate2xo/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	appState, cmds, db := initStateAndCommands()
	defer db.Close()

	command := buildCommandFrom(os.Args)
	err := cmds.run(appState, command)
	if err != nil {
		log.Fatalf("- Error: %v", err)
	}
}

func initStateAndCommands() (*state, commands, *sql.DB) {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("- Error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("- Error initializing database: %v", err)
	}
	dbQueries := database.New(db)
	appState := &state{cfg: &cfg, db: dbQueries}

	cmds := commands{
		registered: make(map[string]func(*state, command) error),
	}
	registerCommands(cmds)

	return appState, cmds, db
}

func registerCommands(cmds commands) {
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
}

func buildCommandFrom(args []string) command {
	if len(args) < 2 {
		log.Fatal("Error: please enter a command name")
	}
	command := command{name: args[1], args: args[2:]}

	return command
}

type state struct {
	cfg *config.Config
	db  *database.Queries
}
