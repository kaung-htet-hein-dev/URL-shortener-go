package main

import (
	"fmt"
	"kaung-htet-hein-dev/URL-shortener-go/handler"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
		fmt.Println("No PORT environment variable detected, defaulting to port 3001")
	}

	e := echo.New()

	e.GET("/", handler.HandleHomePage)

	e.GET("/health", handler.HandleHealthCheck)
	e.GET("/:code", handler.HandleRedirectURL)
	e.POST("/shorten-url", handler.HandleShortenURL)

	e.Logger.Fatal(e.Start(":" + port))
}
