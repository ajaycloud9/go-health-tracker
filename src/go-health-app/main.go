package main

import (
	"fmt"
	"net/http"
)

// CORS middleware to add CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, or replace "*" with specific origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Hello handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Send "Hello, Vijay!" as the HTTP response
	fmt.Fprintf(w, "Hello, Vijay!")
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register the handler for the "/hello" route
	mux.HandleFunc("/hello", helloHandler)

	// Wrap the ServeMux with the CORS middleware
	handlerWithCORS := corsMiddleware(mux)

	// Start the HTTP server on localhost:8332
	fmt.Println("Server is running now on http://localhost:8332/")
	err := http.ListenAndServe(":8332", handlerWithCORS)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
