# URL Shortener (Go)

A simple backend-only URL shortener written in pure Go, with no external dependencies. It takes a long URL, generates a short code for it, and redirects anyone who visits the short link back to the original URL.

## Features

- Generate a short code for any URL
- Redirect from the short code to the original URL
- No external packages — uses only Go's standard library

## Requirements

- [Go](https://go.dev/dl/) 1.20 or newer

## Getting Started

```bash
go run main.go
```

The server starts on `http://localhost:8080`.

## API Endpoints

### Create a short URL
```
POST /shorten
Content-Type: application/json

{
  "url": "https://www.example.com/some/very/long/path"
}
```

**Response:**
```json
{
  "original_url": "https://www.example.com/some/very/long/path",
  "short_url": "http://localhost:8080/xv21MM"
}
```

### Redirect to the original URL
```
GET /{code}
```
Redirects (HTTP 302) to the original URL associated with that code.

## Project Structure

```
.
├── main.go
└── README.md
```

## How It Works

1. A client sends a long URL to `POST /shorten`.
2. The server generates a random 6-character alphanumeric code and stores the mapping `code → URL` in an in-memory map.
3. Visiting `GET /{code}` looks up the code and issues a redirect to the original URL.

## Note

Data is stored in memory only — it resets whenever the server restarts. This is intentional for keeping the code simple and easy to explain.

## License

Free to use for learning and coursework purposes.