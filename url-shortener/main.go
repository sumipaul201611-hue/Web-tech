package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Stores short code -> original URL
var urls = make(map[string]string)

// Request structure
type Request struct {
	URL string `json:"url"`
}

// Generate a 6-character short code
func generateCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano()) // largeNumber

	code := make([]byte, 6)

	for i := 0; i < 6; i++ {
		code[i] = letters[rand.Intn(len(letters))]
	}

	return string(code) // byte->string 
}

// POST /shorten
func shortenHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read JSON
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.URL == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Generate short code
	code := generateCode()

	// Store URL
	urls[code] = req.URL

	// Create short URL
	shortURL := "http://localhost:8080/" + code

	// Send response
	response := map[string]string{
		"original_url": req.URL,
		"short_url":    shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GET /code
func redirectHandler(w http.ResponseWriter, r *http.Request) {

	// Get short code from URL
	code := strings.TrimPrefix(r.URL.Path, "/")

	// Find original URL
	originalURL, exists := urls[code]

	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	// Redirect to original URL
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {

	// /shorten endpoint
	http.HandleFunc("/shorten", shortenHandler)

	// Short URL endpoint
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server running at http://localhost:8080")

	// Start server
	http.ListenAndServe(":8080", nil)
}
