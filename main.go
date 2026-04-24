package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// RequestLogger middleware logs HTTP requests with timing information
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s\n", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Request completed in %v\n", time.Since(start))
	})
}

// DataResponse represents the structure of our data endpoint
type DataResponse struct {
	Message   string        `json:"message"`
	Timestamp string        `json:"timestamp"`
	Items     []DataItem    `json:"items"`
}

type DataItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Value string `json:"value"`
}

func main() {
	// Define HTTP server configuration
	mux := http.NewServeMux()
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: RequestLogger(mux),
	}

	// Register routes
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/data", handleData)

	// Start server
	log.Printf("Starting HTTP server on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v\n", err)
	}
}

// handleRoot handles requests to the root path
func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Welcome to the HTTP server"}`)
}

// handleHealth handles health check requests
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "healthy"}`)
}

// handleData handles requests to the data endpoint with JSON response
func handleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := DataResponse{
		Message:   "Sample data from server",
		Timestamp: time.Now().Format(time.RFC3339),
		Items: []DataItem{
			{ID: 1, Name: "Item One", Value: "Value A"},
			{ID: 2, Name: "Item Two", Value: "Value B"},
			{ID: 3, Name: "Item Three", Value: "Value C"},
		},
	}
	
	json.NewEncoder(w).Encode(response)
}
