package main

import (
	"blog/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("HELLLO FROM MAIN")
	configStruct := &config.Config{}

	stateInstance := &config.State{
		Config: configStruct,
	}

	commands := &config.Commands{
		Handlers: make(map[string]func(*config.State, config.Command) error),
	}

	commands.Register("login", config.HandlerLogin)
	if len(os.Args) < 2 {
		log.Fatalf("Command name is required")
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	cmd := config.Command{
		Name: cmdName,
		Args: args,
	}

	err := commands.Run(stateInstance, cmd)
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}

	configStruct.Read()
}
