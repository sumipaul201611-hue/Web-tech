package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status": "ok",
		"code":   200,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)

	fmt.Println("Server started on 8080. Open http://localhost:8080/health")
	http.ListenAndServe(":8080", nil)
}


