# URL Shortener (Go)

A simple URL shortener web service built with Go and Echo framework. It allows users to submit long URLs and receive shortened links, which redirect to the original URLs. Includes a minimal web interface and stores data using a database.

## Initial Setup

1. Install Go (https://golang.org/doc/install)
2. Clone this repository
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Prepare the database (SQLite by default):
   - The app will create `shortener.db` automatically on first run.
   - No manual migration needed for basic usage.

## Features

- Shorten any URL
- Redirect using short codes
- Health check endpoint
- Minimal HTML frontend

## Usage

1. Run the server: `go run main.go`
2. Open `http://localhost:3001` in your browser
3. Enter a URL to get a shortened link

## Endpoints

- `/shorten-url` (POST): Shorten a URL
- `/:code` (GET): Redirect to original URL
- `/health` (GET): Health check

---

Built with Go, Echo, and GORM.
