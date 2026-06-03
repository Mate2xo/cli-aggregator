package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}
type commands struct {
	registered map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, found := c.registered[cmd.name]
	if !found {
		return fmt.Errorf("handler for command '%v' not found", cmd.name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registered[name] = f
}
