package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// =========================
// HEALTH RESPONSE
// =========================

type healthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

// =========================
// LOGIN REQUEST
// =========================

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// =========================
// LOGIN RESPONSE
// =========================

type loginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// =========================
// HEALTH HANDLER
// =========================

func healthHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := healthResponse{
		Status: "ok",
		Time:   time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

// =========================
// LOGIN HANDLER
// =========================

func loginHandler(w http.ResponseWriter, r *http.Request) {

	// Step 2:
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Step 3 + Step 4:
	// Create an empty LoginRequest
	var req LoginRequest

	// Read JSON from request body
	// and put the data inside req
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// Step 5:
	// Hardcoded credentials
	const correctUsername = "admin"
	const correctPassword = "password123"

	// Check username and password
	if req.Username != correctUsername ||
		req.Password != correctPassword {

		http.Error(
			w,
			"invalid username or password",
			http.StatusUnauthorized,
		)
		return
	}

	// Step 6:
	// Successful login response
	resp := loginResponse{
		Status:  "ok",
		Message: "login successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(resp)
}

// =========================
// MAIN
// =========================

func main() {

	// Create router
	mux := http.NewServeMux()

	// Register endpoints
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/login", loginHandler)

	// Server address
	addr := ":8080"

	log.Printf("Server starting on %s ...", addr)

	// Start server
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}