package main

import (
	"blog/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFeedList(s *state, cmd command) error {
  feedData, err := s.db.ListFeeds(context.Background())
  if err != nil {
    return err
  }
  for _, feed := range feedData {
    fmt.Printf("Feed Name: %s\n",feed.Name)
    fmt.Printf("Url: %s\n", feed.Url)
    fmt.Printf("User Name: %s\n", feed.Name_2)
  }
  return nil
}


func handlerFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("There need to be at least 2 args")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Println("feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=============================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
