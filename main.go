package main

import (
	"blog/internal/config"
	"blog/internal/database"
	"database/sql"
	"context"
  "fmt"
	"log"
	_ "log"
	"os"
	_ "os"

	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

const dbURL = "postgres://mariglenpoleshi:@localhost:5432/gator?sslmode=disable"

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("couldn't read config file")
	}

	// stateInstance.config.DbURL = dbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		config: &cfg,
		db:     dbQueries,
	}

	cmds := commands{
		Handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handlerUsersList)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerFeed))
  cmds.register("feeds", handlerFeedList)
  cmds.register("follow", middlewareLoggedIn(handlerFeedFollow))
  cmds.register("following", middlewareLoggedIn(handlerFeedFollowsForUser))
  cmds.register("unfollow", middlewareLoggedIn(handlerFeedUnfollow))
  cmds.register("browse", middlewareLoggedIn(handleBrowse))


	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...] ")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
  return func(s *state, cmd command) error {
    user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
    if err != nil {
      return err 
    }
    return handler(s, cmd, user)
  }
}
