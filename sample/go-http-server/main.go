package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Send "Hello, World!" as the HTTP response
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Register the handler for the "/" route
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server on localhost:8080
	fmt.Println("Server is running on http://localhost:8332/")
	err := http.ListenAndServe(":8332", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
