package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

type Limit struct {
	count int
	start time.Time
}

var (
	users = map[string]*Limit{}
	mu    sync.Mutex //Mutual Exclusion
)

func rateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //127.0.0.1:54321
		ip, _, _ := net.SplitHostPort(r.RemoteAddr) //ip->host->port

		mu.Lock()

		u := users[ip]

		if u == nil || time.Since(u.start) >= time.Minute {
			users[ip] = &Limit{1, time.Now()}
			u = users[ip]
		} else {
			u.count++
		}

		mu.Unlock()

		if u.count > 5 {
			http.Error(w, "Too Many Requests", 429)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	http.Handle("/", rateLimiter(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello!")
		}),
	))

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}

