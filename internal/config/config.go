package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const fileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (config *Config) SetUser(newUsername string) error {
	config.CurrentUserName = newUsername
	return write(*config)
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonData, err := os.Open(fullPath)
	if err != nil {
		return Config{}, fmt.Errorf("")
	}
	defer jsonData.Close()

	var config Config
	decoder := json.NewDecoder(jsonData)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	// fmt.Printf("dbURL: %s\n,  username: %s\n", config.DbURL, config.CurrentUserName)
	return config, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homePath, fileName)
	return fullPath, nil
}

func write(config Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return err
	}

	return nil
}
