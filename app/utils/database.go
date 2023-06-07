package utils

import (
	"log"

	"github.com/sorasora46/Tungleua-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := config.GetDatabaseDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect to Database error: %v", err)
	}

	DB = db

	return nil
}
