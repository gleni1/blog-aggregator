package main

import (
	"context"
	"fmt"
	"time"

	"blog/internal/database"

	"github.com/google/uuid"
)

func handlerUsersList(s *state, cmd command) error {
	users, err := s.db.Getusers(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching users from the database.")
	}
	for _, user := range users {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	return nil
}

func handleReset(s *state, cmd command) error {
	err := s.db.ClearData(context.Background())
	if err != nil {
		return fmt.Errorf("Data could not be cleared: %w", err)
	}

	fmt.Println("Success deleting all data from the users table")
	return nil
}

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
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
