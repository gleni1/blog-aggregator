package config

import (
  "fmt"
)

type Config struct {
  DbURL           string `json:"db_url"`
  CurrentUserName string `json:"current_user_name"` 
}

func DoSth () {
  fmt.Println("Hello from the inner package")
}
