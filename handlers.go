package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("The login command expects the username argument")
	}

	username := cmd.arguments[0]

	s.conf.SetUser(username)

	fmt.Println("Username was successfully set to:", username)

	return nil
}
