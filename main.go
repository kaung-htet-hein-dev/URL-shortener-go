package main

import (
	"fmt"
	"kaung-htet-hein-dev/URL-shortener-go/db"
	"kaung-htet-hein-dev/URL-shortener-go/handler"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func registerRoutes(e *echo.Echo, h *handler.Handler) {
	e.File("/", "index.html")
	e.GET("/health-check", h.HandleHealthCheck)
	e.GET("/:code", h.HandleRedirectURL)
	e.POST("/shorten-url", h.HandleShortenURL)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
		fmt.Println("No PORT environment variable detected, defaulting to port 3001")
	}

	e := echo.New()

	var database *gorm.DB = db.DB
	h := handler.NewHandler(database)
	registerRoutes(e, h)

	e.Logger.Fatal(e.Start(":" + port))
}
