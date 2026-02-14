package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rQxwX3/gator/internal/database"
	"github.com/rQxwX3/gator/internal/rss"
	"os"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("The login command requires username argument")
	}

	name := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		os.Exit(1)
	}

	s.conf.SetUser(name)

	fmt.Println("Username was successfully set to:", name)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("The register command requires username argument")
	}

	name := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != sql.ErrNoRows {
		os.Exit(1)
	}

	currentTime := time.Now()

	user, err := s.db.CreateUser(context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
			Name:      name,
		},
	)

	if err != nil {
		return nil
	}

	s.conf.SetUser(user.Name)
	fmt.Println("Successfully created user:", user)

	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("The reset command does not take any arguments")
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("The users command does not take any arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Print("* ", user.Name)

		if user.Name == s.conf.CurrentUserName {
			fmt.Print(" (current)")
		}

		fmt.Print("\n")
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("The agg command does not take any arguments")
	}

	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return errors.New("The addfeed command requires name and URL arguments")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return nil
	}

	currentTime := time.Now()

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    currentUser.ID,
	})

	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
