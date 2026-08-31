package main

import (
	"fmt"
	"net/http"
)

func sum(w http.ResponseWriter, r *http.Request) {
	a := 50
	b := 100

	sum := a + b
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"sum": %d}`, sum)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{
	"status": "okay","status_code":200}`)
}

func main() {
	http.HandleFunc("/sum", sum)
	http.HandleFunc("/health", health)

	fmt.Println("=================================")
	fmt.Println("Server started successfully!")
	fmt.Println("Sum API:    http://localhost:8080/sum")
	fmt.Println("Health API: http://localhost:8080/health")
	fmt.Println("=================================")

	http.ListenAndServe(":8080", nil)

}
