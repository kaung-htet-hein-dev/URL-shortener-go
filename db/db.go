package db

import (
	"kaung-htet-hein-dev/URL-shortener-go/entity"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(sqlite.Open("shortener.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&entity.URL{})

	DB = db
}
