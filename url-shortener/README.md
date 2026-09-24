URL Shortener (Go)

A simple backend-only URL shortener written in pure Go, with no external dependencies. It takes a long URL, generates a short code for it, and redirects anyone who visits the short link back to the original URL.

Features
Generate a short code for any valid URL
Redirect from the short code to the original URL
List all shortened URLs with click counts
Data persists across restarts (saved to urls.json)
No external packages — uses only Go's standard library
Requirements
Go 1.20 or newer
Getting Started

Clone the repo and run:

bash
go run main.go

The server starts on http://localhost:8080.

API Endpoints
Create a short URL
POST /shorten
Content-Type: application/json

{
  "url": "https://www.example.com/some/very/long/path"
}

Response:

json
{
  "code": "xv21MM",
  "short_url": "http://localhost:8080/xv21MM",
  "original_url": "https://www.example.com/some/very/long/path"
}
Redirect to the original URL
GET /{code}

Redirects (HTTP 302) to the original URL associated with that code.

List all shortened URLs
GET /urls

Response:

json
[
  {
    "code": "xv21MM",
    "original_url": "https://www.example.com/some/very/long/path",
    "created_at": "2026-09-24T10:00:00Z",
    "clicks": 3
  }
]
Health check
GET /health
Project Structure
.
├── main.go       # server, routes, and storage logic
├── urls.json     # auto-generated data file (created on first run)
└── README.md
How It Works
A client sends a long URL to POST /shorten.
The server generates a random 6-character alphanumeric code and stores the mapping code → URL in memory, then saves it to urls.json.
Visiting GET /{code} looks up the code and issues a redirect to the original URL.
A mutex guards the in-memory store so concurrent requests don't corrupt it.
License

Free to use for learning and coursework purposes