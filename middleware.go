package main

import (
	"context"
)

func middlewareLoggedIn(handler loggedInHandler) handler {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
