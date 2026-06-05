package main

import (
	"log"
	"os"

	"github.com/Mate2xo/gator/internal/config"
)

func main() {
	appState, cmds := initStateAndCommands()
	command := buildCommandFrom(os.Args)

	err := cmds.run(appState, command)
	if err != nil {
		log.Fatalf("- Error: %v", err)
	}
}

func initStateAndCommands() (*state, commands) {
	cfg, err := config.Read()
	appState := &state{cfg: &cfg}
	if err != nil {
		log.Fatalf("- Error reading config: %v", err)
	}

	cmds := commands{
		registered: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	return appState, cmds
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
}
