package util

import (
	"fmt"
	"math/rand"

	"kaung-htet-hein-dev/URL-shortener-go/db"
	"kaung-htet-hein-dev/URL-shortener-go/entity"

	"github.com/labstack/echo/v4"
)

func GenerateRandomCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result = make([]byte, 8)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func GetHostDomainAddress(c echo.Context) string {
	host := c.Request().Host
	scheme := "http"

	if c.Request().TLS != nil {
		scheme = "https"
	}

	return fmt.Sprintf("%s://%s", scheme, host)
}

func FindInDB(code string) (string, error) {
	var url entity.URL

	err := db.DB.Where("shortened_code = ?", code).First(&url).Error

	if err != nil {
		return "", err
	}

	return url.OriginalURL, nil
}
