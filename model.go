package main

import "gorm.io/gorm"

type URL struct {
	gorm.Model

	ShortenedCode string `json:"shortened_code"`
	OriginalURL   string `json:"original_url"`
}

type Request struct {
	OriginalURL string `json:"original_url"`
}
