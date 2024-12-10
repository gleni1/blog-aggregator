package main

import (
	"context"
	"fmt"
	"time"

	"blog/internal/database"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("There is no arguments provided")
	}

	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("You cannot login to an account that does not exist.")
	}

	err = s.config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error setting user to the provided username")
	}
	fmt.Println("success setting new username")
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("A name needs to be passed as an arg")
	}
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	}
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Error while setting user: %w", err)
	}

	fmt.Printf("Success! User was created.\nUser ID: %s\nName: %s\n",
		user.ID, user.Name)

	return nil
}
