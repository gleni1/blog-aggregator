package main

import (
  "context"
  "fmt"
  "strconv"

	"blog/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
  limit := 2
  if len(cmd.Args) == 1 {
    if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err != nil {
      limit = specifiedLimit
    } else {
      fmt.Printf("invalid limit: %w", err)
    }
  }
  postsForUserParams := database.GetPostsForUserParams {
    UserID:   user.ID,
    Limit:    int32(limit),
  }
  posts, err := s.db.GetPostsForUser(context.Background(), postsForUserParams)
  if err != nil {
    return fmt.Errorf("Error getting posts for user: %s", user.Name)
  }
  for _, post := range posts {
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
  return nil

}
