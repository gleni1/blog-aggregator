package main

import (
  "fmt"
  "blog/internal/config"
  "log"
)

func main() {
  if err := run(); err != nil {
    log.Fatalf("Application error: %w", err)
  }
}

func run () error {
  c, err := config.Read()
  if err != nil {
    log.Fatalf("Error reading config: %v", err)
  }
  fmt.Printf("Config: %+v\n",c)

  newUserName := "Glen"
  err = c.SetUser(newUserName)
  if err != nil {
    log.Fatalf("Error setting new username: %v", err)
  }

  fmt.Printf("Config: %+v\n", c)
  return nil
}
