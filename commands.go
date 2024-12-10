package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	Handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("Error running given command")
	}
	return f(s, cmd)
}
