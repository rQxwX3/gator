package main

import (
	"github.com/rQxwX3/gator/internal/config"
)

type state struct {
	conf *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdHandler := c.cmdMap[cmd.name]

	if err := cmdHandler(s, cmd); err != nil {
		return err
	}

	return nil
}



// import (
// 	"errors"
// 	"fmt"
// 	"github.com/rQxwX3/gator/internal/types"
// )
//
// func handlerLogin(s *types.State, cmd types.Command) error {
// 	if len(cmd.Arguments) != 1 {
// 		return errors.New("The login command expects the username argument")
// 	}
//
// 	username := cmd.Arguments[0]
//
// 	s.Conf.SetUser(username)
//
// 	fmt.Println("Username was successfully set to:", username)
//
// 	return nil
// }
