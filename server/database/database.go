package database

import (
	"os"
	"fmt"

	"github.com/donny-c-1/sorcemoola/server/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func Connect () {
	connStr := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config {})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established successfully.")

	return nil
}

func Migrate () {
	if DB == nil {
		return fmt.Errorf("database connection not established")
	}

	err := DB.AutoMigrate(&models.User {})
	if err != nil {
		return fmt.Errorf("Failed to run migrations: %w", err)
	}
	log.Println("Database migrations comoleted successfully")
	
	return nil
}