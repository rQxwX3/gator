package main

import (
	"errors"
	"github.com/rQxwX3/gator/internal/config"
	"github.com/rQxwX3/gator/internal/database"
)

type state struct {
	conf *config.Config
	db   *database.Queries
}

type command struct {
	name      string
	arguments []string
}

type handler func(*state, command) error
type loggedInHandler func(*state, command, database.User) error
type cmdmap map[string]handler

type commands struct {
	cmdmap cmdmap
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.cmdmap[cmd.name]
	if !ok {
		return errors.New(cmd.name + " is not a registered command")
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f handler) error {
	if _, ok := c.cmdmap[name]; !ok {
		c.cmdmap[name] = f
	}

	return nil
}
