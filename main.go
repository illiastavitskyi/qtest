package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define HTTP server configuration
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	// Register routes
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)

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
