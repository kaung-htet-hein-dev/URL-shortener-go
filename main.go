package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
		panic("failed to connect database")
	}

	db.AutoMigrate(&URL{})
	db.Exec("ALTER TABLE urls DROP COLUMN shortened_url;")

	DB = db
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
		fmt.Println("No PORT environment variable detected, defaulting to port 3001")
	}

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, this is URL Shortener")
	})
	e.GET("/:code", handleRedirectURL)
	e.POST("/shorten-url", handleShortenURL)

	e.Logger.Fatal(e.Start(":" + port))
}
