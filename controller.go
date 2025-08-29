package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleShortenURL(c echo.Context) error {
	req := new(Request)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if req.OriginalURL == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Please provide url."})
	}

	randomCode := GenerateRandomCode()

	data := DB.Create(&URL{
		OriginalURL:   req.OriginalURL,
		ShortenedCode: randomCode,
	})

	if data.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": data.Error.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"shortened_url": GetHostDomainAddress(c) + "/" + randomCode,
		"original_url":  req.OriginalURL,
	})
}

func handleRedirectURL(c echo.Context) error {
	code := c.Param("code")
	completedURL, err := FindInDB(code)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "URL not found"})
	}

	return c.Redirect(http.StatusMovedPermanently, completedURL)
}
