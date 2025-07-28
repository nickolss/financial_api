package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	dbURL := "postgres://user:pass@localhost:5432/database"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("error to connect to database: %s", err)
	}

	DB = db
	return nil
}