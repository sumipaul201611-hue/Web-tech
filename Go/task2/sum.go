package main

import (
	"fmt"
	"net/http"
)


func sum(w http.ResponseWriter, r *http.Request) {
	a:=50
	b:=50
	sum := a + b
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"sum": %d}`, sum)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{
    "status": "okay","status_code": 200}`)
}

func main() {
	http.HandleFunc("/sum", sum)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)

}