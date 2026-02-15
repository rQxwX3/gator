package main

import (
	"context"
	"github.com/rQxwX3/gator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.User) error) cmdhandler {

	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
