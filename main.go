package main

import (
	"fmt"
	"kaung-htet-hein-dev/URL-shortener-go/handler"
	"net/http"
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

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, this is URL Shortener")
	})
	e.GET("/:code", handler.HandleRedirectURL)
	e.POST("/shorten-url", handler.HandleShortenURL)

	e.Logger.Fatal(e.Start(":" + port))
}
