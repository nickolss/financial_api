package config

import (
	"fmt"

	"github.com/nickolss/financial_api/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	dbURL := Env.Database_url
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("error to connect to database: %s", err)
	}

	DB = db
	RealizeMigrations()
	return nil
}

func RealizeMigrations() {
	DB.AutoMigrate(&schemas.Category{}, &schemas.Client{}, &schemas.FinancialEvent{})
}
