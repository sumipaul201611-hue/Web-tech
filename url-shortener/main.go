// URL Shortener
// How it works:
//   1. Client sends POST /shorten with JSON {"url": "https://example.com/very/long/link"}
//   2. Server generates a random 6-character code (like "aZ3kLp")
//   3. Server stores code -> original URL in memory (and saves to a JSON file so data survives restarts)
//   4. Client can now visit GET /aZ3kLp and the server redirects them to the original long URL
//
// Endpoints:
//   POST /shorten      -> create a short URL
//   GET  /{code}        -> redirect to the original URL
//   GET  /urls          -> list every short URL created so far (handy for the demo)
//   GET  /health         -> simple health check
 
package main
 
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)
 
// URLEntry represents one shortened link.
type URLEntry struct {
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	Clicks      int       `json:"clicks"`
}
 
// Store keeps everything in memory, protected by a mutex so concurrent
// requests don't corrupt the map. It also persists to a JSON file.
type Store struct {
	mu   sync.Mutex
	data map[string]*URLEntry
	file string
}
 
func NewStore(file string) *Store {
	s := &Store{data: make(map[string]*URLEntry), file: file}
	s.load()
	return s
}
 
// load reads existing short URLs from disk (if the file exists) so the
// server "remembers" links even after being restarted.
func (s *Store) load() {
	bytes, err := os.ReadFile(s.file)
	if err != nil {
		return // no file yet, start fresh
	}
	var entries []*URLEntry
	if err := json.Unmarshal(bytes, &entries); err != nil {
		return
	}
	for _, e := range entries {
		s.data[e.Code] = e
	}
}
 
// save writes the current map to disk as JSON.
func (s *Store) save() {
	entries := make([]*URLEntry, 0, len(s.data))
	for _, e := range s.data {
		entries = append(entries, e)
	}
	bytes, _ := json.MarshalIndent(entries, "", "  ")
	_ = os.WriteFile(s.file, bytes, 0644)
}
 
// generateCode creates a random 6-character base62 string, e.g. "aZ3kLp".
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
 
func generateCode(n int) string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}
 
// Add stores a new URL and returns the short code it was given.
// If the same URL was already shortened before, it reuses the old code.
func (s *Store) Add(originalURL string) *URLEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
 
	for _, e := range s.data {
		if e.OriginalURL == originalURL {
			return e // avoid duplicate codes for the same link
		}
	}
 
	var code string
	for {
		code = generateCode(6)
		if _, exists := s.data[code]; !exists {
			break
		}
	}
 
	entry := &URLEntry{Code: code, OriginalURL: originalURL, CreatedAt: time.Now()}
	s.data[code] = entry
	s.save()
	return entry
}
 
// Get looks up the original URL for a code and counts the click.
func (s *Store) Get(code string) (*URLEntry, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.data[code]
	if ok {
		e.Clicks++
		s.save()
	}
	return e, ok
}
 
// All returns every stored entry, newest first.
func (s *Store) All() []*URLEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	entries := make([]*URLEntry, 0, len(s.data))
	for _, e := range s.data {
		entries = append(entries, e)
	}
	return entries
}
 
var store *Store
 
// --- HTTP Handlers ---
 
type shortenRequest struct {
	URL string `json:"url"`
}
 
type shortenResponse struct {
	Code        string `json:"code"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
 
func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "use POST", http.StatusMethodNotAllowed)
		return
	}
 
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid JSON body"}`, http.StatusBadRequest)
		return
	}
	req.URL = strings.TrimSpace(req.URL)
	if req.URL == "" {
		http.Error(w, `{"error":"url is required"}`, http.StatusBadRequest)
		return
	}
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		http.Error(w, `{"error":"url must start with http:// or https://"}`, http.StatusBadRequest)
		return
	}
 
	entry := store.Add(req.URL)
	resp := shortenResponse{
		Code:        entry.Code,
		ShortURL:    fmt.Sprintf("http://localhost:8080/%s", entry.Code),
		OriginalURL: entry.OriginalURL,
	}
 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
 
func handleRedirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	entry, ok := store.Get(code)
	if !ok {
		http.Error(w, "short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, entry.OriginalURL, http.StatusFound) // 302 redirect
}
 
func handleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.All())
}
 
func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK - URL shortener is running")
}
 
func main() {
	store = NewStore("urls.json")
 
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handleShorten)
	mux.HandleFunc("/urls", handleList)
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/", handleRedirect) // catches any /{code} request
 
	fmt.Println("URL Shortener running on http://localhost:8080")
	fmt.Println("POST   /shorten   { \"url\": \"https://...\" }")
	fmt.Println("GET    /{code}    redirects to original URL")
	fmt.Println("GET    /urls      lists all shortened URLs")
	log.Fatal(http.ListenAndServe(":8080", mux))
}