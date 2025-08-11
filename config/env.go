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

	rawPort := os.Getenv("PORT")
	dbURL := os.Getenv("PSQL_URL")

	if rawPort == "" {
		return fmt.Errorf("missing required environment variable: PORT")
	}
	if dbURL == "" {
		return fmt.Errorf("missing required environment variable: PSQL_URL")
	}

	Env = Environment{
		Port:         fmt.Sprintf(":%s", rawPort),
		Database_url: dbURL,
	}

	return nil
}
