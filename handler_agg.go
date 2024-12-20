package main

import (
	"context"
	"fmt"
  "time"
  "log"
	"blog/internal/database"
	"github.com/google/uuid"
	"database/sql"
)

func handlerAgg(s *state, cmd command) error {
  if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
    return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
  }
  timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
  if err != nil {
    return fmt.Errorf("Invalid duration: %w", err)
  }
  log.Printf("Collecting feeds every %s...", timeBetweenRequests)

  ticker := time.NewTicker(timeBetweenRequests)

  for ; ; <-ticker.C {
    scrapeFeeds(s)
  }
}

func scrapeFeeds(s *state) {
  feed, err := s.db.GetNextFeedToFetch(context.Background())
  if err != nil {
    log.Println("Couldn't get next feeds to fetch", err)
    return
  }
  log.Println("Found a feed to fetch!")
  scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
  _, err := db.MarkFeedFetched(context.Background(), feed.ID)
  if err != nil {
    log.Printf("Couldn't mark feed %s as fetched: %v", feed.Name, err)
    return
  }
  feedData, err := fetchFeed(context.Background(), feed.Url)
  if err != nil {
    log.Printf("Couldn't collect feed: %w", err)
    return
  }
  for _, item := range feedData.Channel.Item {
    //fmt.Printf("Found post: %s\n", item.Title)
    err = storePosts(db, item, feed.ID)
    if err != nil {
      log.Printf("Error storing post: %w", err)
    }
  }
  log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}


func storePosts(db *database.Queries, item RSSItem, feedID uuid.UUID) error {
  layout := "Mon, 02 Jan 2006 15:04:05 -0700"
  parsedTime, err := time.Parse(layout, item.PubDate)
  if err != nil {
    return fmt.Errorf("Error parsing time in the right format: %w", err)
  }
  postParams := database.CreatePostParams {
    ID:           uuid.New(),
    CreatedAt:    time.Now().UTC(),
    UpdatedAt:    time.Now().UTC(),
    Title:        item.Title,
    Url:          sql.NullString{
      String:   item.Link,
      Valid:    true,
    },
    Description:  item.Description,
    PublishedAt:  parsedTime, 
    FeedID:       feedID, 
  }
  post, err := db.CreatePost(context.Background(), postParams)
  if err != nil {
    return fmt.Errorf("Error storing post: %w", err)
  }
  fmt.Printf("Success storing post with title: %s\n", post.Title)
  return nil
}
