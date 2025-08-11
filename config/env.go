package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Database_url string
	Port         string
}

var Env Environment

func InitEnv() error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, continuing with environment variables")
	}

	Env = Environment{
		Port:         os.Getenv("PORT"),
		Database_url: os.Getenv("PSQL_URL"),
	}

	if Env.Port == "" {
		return fmt.Errorf("missing required environment variable: PORT")
	}
	if Env.Database_url == "" {
		return fmt.Errorf("missing required environment variable: PSQL_URL")
	}

	return nil
}
