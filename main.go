package main

import (
	"log"
	"os"

	"github.com/Mate2xo/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("- Error reading config: %v", err)
	}

	appState := &state{cfg: &cfg}
	cmds := commands{
		registered: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Error: please enter a command name")
	}
	command := command{name: args[1], args: args[2:]}

	err = cmds.run(appState, command)
	if err != nil {
		log.Fatalf("- Error: %v", err)
	}
}

type state struct {
	cfg *config.Config
}
