package main

import (
	"blog/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFeedFollowsForUser(s *state, cmd command, user database.User) error {
  // Get user data from the databse based on the name provided in args 
  userName := s.config.CurrentUserName 
  userId := user.ID
  
  // Get follows for user based on user ID
  feedFollowerInfoList, err := s.db.GetFeedFollowsForUser(context.Background(), userId)
  if err != nil {
    return fmt.Errorf("Error getting the feed follows for user")
  }

  // Print user name and all feeds they follow
  fmt.Printf("User Name: %s\n", userName)
  for _, followRow := range feedFollowerInfoList {
    fmt.Printf("Feed Name: %s\n", followRow.FeedName)
  }
  return nil
}

func handlerFeedFollow(s *state, cmd command, user database.User) error {
  if len(cmd.Args) != 1 {
    return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
  }

  url := cmd.Args[0]
  feedByUrlRow, err := s.db.GetFeedByURL(context.Background(), url)
  if err != nil {
    return fmt.Errorf("Error retrieving feed row by url")
  }

  feedId := feedByUrlRow.ID

  params := database.CreateFeedFollowParams {
    ID:           uuid.New(),
    CreatedAt:    time.Now(),
    UpdatedAt:    time.Now(),
    UserID:       user.ID,
    FeedID:       feedId,       
  }
  recordRow, err := s.db.CreateFeedFollow(context.Background(), params)
  if err != nil {
    fmt.Errorf("Error creating feed follow record")
  }
  fmt.Printf("Success creating feed follow record\n")
  fmt.Printf("Feed Name: %s\n", recordRow.FeedName)
  fmt.Printf("Current User: %s\n", user.Name)
  return nil
}


func handlerFeedList(s *state, cmd command) error {
  feedData, err := s.db.ListFeeds(context.Background())
  if err != nil {
    return err
  }

  if len(feedData) == 0 {
    fmt.Println("No feeds found.")
    return nil
  }

  for _, feed := range feedData {
    fmt.Printf("Feed Name: %s\n",feed.Name)
    fmt.Printf("Url: %s\n", feed.Url)
    fmt.Printf("User Name: %s\n", feed.Name_2)
  }
  return nil
}


func handlerFeed(s *state, cmd command, user database.User) error {
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

  feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
    ID:         uuid.New(),
    CreatedAt:  time.Now().UTC(),
    UpdatedAt:  time.Now().UTC(),
    UserID:     user.ID,
    FeedID:     feed.ID,
  })
  if err != nil {
    return fmt.Errorf("Couldn't create feed follow: %w", err)
  }

	fmt.Println("feed created successfully:")
	printFeed(feed)
	fmt.Println()
  fmt.Println("Feed followed successfully")
  fmt.Printf("Username: %s\n", feedFollow.UserName)
  fmt.Printf("FeedName: %s\n", feedFollow.FeedName)
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
