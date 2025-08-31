package handler

import (
	"kaung-htet-hein-dev/URL-shortener-go/entity"
	"kaung-htet-hein-dev/URL-shortener-go/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the URL Shortener")
}

func (h *Handler) HandleShortenURL(c echo.Context) error {
	req := new(entity.Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if req.OriginalURL == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Please provide url."})
	}
	randomCode := util.GenerateRandomCode()
	data := h.DB.Create(&entity.URL{
		OriginalURL:   req.OriginalURL,
		ShortenedCode: randomCode,
	})
	if data.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": data.Error.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"shortened_url": util.GetHostDomainAddress(c) + "/" + randomCode,
		"original_url":  req.OriginalURL,
	})
}

func (h *Handler) HandleRedirectURL(c echo.Context) error {
	code := c.Param("code")
	completedURL, err := util.FindInDB(code)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "URL not found"})
	}
	return c.Redirect(http.StatusMovedPermanently, completedURL)
}
