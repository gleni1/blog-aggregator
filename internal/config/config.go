package config

import (
  "os"
  "fmt"
  "log"
  "encoding/json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type State struct {
  Config *Config
}

type Command struct {
  Name      string
  Args      []string
}

type Commands struct {
  Handlers map[string]func(*State, Command)error
}

func (commands *Commands) Run(state *State, cmd Command) error {
  handler, ok := commands.Handlers[cmd.Name]
  if !ok {
    return fmt.Errorf("Command not found")
  }
  return handler(state, cmd)
}

func (commands *Commands) Register(name string, f func(*State, Command) error) {
  if commands.Handlers == nil {
    commands.Handlers = make(map[string]func(*State, Command) error)
  }
  commands.Handlers[name] = f
}

func HandlerLogin(s *State, cmd Command) error {
  //userName := "mariglen"
  if len(cmd.Args) == 0 {
    return fmt.Errorf("login handler expects at least one argument: the username")
  }
  err := s.Config.SetUser(cmd.Args[0])
  if err != nil {
    return fmt.Errorf("error setting the username field")
  }
  fmt.Printf("User has been successfully set")
  return nil
}

func (c *Config) Read()  error {
  // put together the path of the file we need
  fileName := ".gatorconfig.json"
  homePath, err := os.UserHomeDir()
  if err != nil {
    fmt.Println("Could not fetch userhomedir")
  }
  filePath := fmt.Sprintf("%s/%s", homePath, fileName)
 
  // open and read the file
  file, err := os.Open(filePath)
  if err != nil {
    log.Fatalf("Error opening file: %v\n", err)
  }
  defer file.Close()

  // store the content of the json file in a Config struct
  decoder := json.NewDecoder(file)
  err = decoder.Decode(&c) 
  if err != nil {
    log.Fatalf("Error decoding the file: %v\n", err) 
  }
  fmt.Printf("dbURL: %s\n, username: %s\n", c.DbURL, c.CurrentUserName)
  return nil
}


func (c *Config) SetUser(newUsername string) error {
  c.CurrentUserName = newUsername

  jsonData, err := json.Marshal(c) 
  if err != nil {
    log.Fatalf("JSON marshaling error: %v", err)
  }

  filePath := "/Users/mariglenpoleshi/.gatorconfig.json" 
  
  err = os.WriteFile(filePath, jsonData, 0644)
  if err != nil {
    fmt.Println("Error writing to file:", err)
    return err
  }
  return nil
}

