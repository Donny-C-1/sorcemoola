package database

import (
	"fmt"
	"log"
	"os"

	"github.com/donny-c-1/sorcemoola/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
	connStr := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established successfully.")

	return nil
}

func Migrate() error {
	if DB == nil {
		return fmt.Errorf("database connection not established")
	}

	err := DB.AutoMigrate(
		&models.User{},
		&models.Campaign{},
		&models.Contribution{},
	)
	if err != nil {
		return fmt.Errorf("Failed to run migrations: %w", err)
	}
	log.Println("Database migrations comoleted successfully")

	return nil
}
