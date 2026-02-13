package main

import (
	"errors"
	"github.com/rQxwX3/gator/internal/config"
)

type state struct {
	conf *config.Config
}

type command struct {
	name      string
	arguments []string
}

type cmdhandler func(*state, command) error
type cmdmap map[string]cmdhandler

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

func (c *commands) register(name string, f cmdhandler) error {
	if _, ok := c.cmdmap[name]; !ok {
		c.cmdmap[name] = f
	}

	return nil
}
