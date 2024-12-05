package config

import (
  "fmt"
  "encoding/json"
  "path/filepath"
  "log"
  "os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
  DbURL           string `json:"db_url"`
  CurrentUserName string `json:"current_user_name"` 
}

func Read() (Config, error) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    return Config{}, fmt.Errorf("error getting home directory")
  }

  pathName := fmt.Sprintf("%s/%s", homeDir, configFileName)

  file, err := os.Open(pathName)
  if err != nil {
    log.Fatalf("Error opening file: %v", err)
  }
  defer file.Close()

  var config Config 
  decoder := json.NewDecoder(file)
  err = decoder.Decode(&config)
  if err != nil {
    return Config{}, fmt.Errorf("error decoding JSON: %v", err)
  }

  return config, nil 
}

func (config *Config) SetUser(userName string) error {
  config.CurrentUserName = userName
  jsonData, err := json.Marshal(config)
  if err != nil {
    return fmt.Errorf("error marshaling config to JSON: %w", err)
  }

  pathName, err := getConfigFilePath(configFileName)
  if err != nil {
    return fmt.Errorf("error getting config filepath: %w", err)
  }

  err = os.WriteFile(pathName, jsonData, 0644)
  if err != nil {
    return fmt.Errorf("error writing to the config file: %w", err)
  }
  return nil
}

func getConfigFilePath(configFileName string) (string ,error) {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    return "", fmt.Errorf("error getting home directory: %w", err)
  }
  pathName := filepath.Join(homeDir, configFileName)
  return pathName, nil
}

func write(cfg Config) error {
  return nil  
}

