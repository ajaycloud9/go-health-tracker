package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ajaycloud9/go-health-tracker/src/go-health-app/domain/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func databaseSetUp() {
	// Define the SQL Server connection string
	// Update with your actual username, password, and database name
	dsn := "root:mysecretpassword@tcp(localhost:3307)/your_database_name?charset=utf8mb4&parseTime=True&loc=Local"
	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to SQL Server successfully!")

	// Auto-migrate the models
	err = db.AutoMigrate(&models.User{}, &models.WeightEntry{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	fmt.Println("Database migrated successfully!")

	// Example: Add a user and associated weight entries
	user := models.User{
		Username: "JohnDoe",
		Email:    "johndoe@example.com",
		WeightEntries: []models.WeightEntry{
			{EntryDate: time.Now(), Weight: 70.5, Comment: "Feeling good!"},
			{EntryDate: time.Now().AddDate(0, 0, -1), Weight: 71.0, Comment: "Yesterday's entry"},
		},
	}

	// Save the user along with their weight entries
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Println("User and weight entries added successfully!")
}

func main() {
	// Create a new ServeMux
	// mux := http.NewServeMux()

	fmt.Println("Setting up the database server")
	databaseSetUp()

	// // Register the handler for the "/hello" route
	// mux.HandleFunc("/hello", helloHandler)

	// // Wrap the ServeMux with the CORS middleware
	// handlerWithCORS := corsMiddleware(mux)

	// // Start the HTTP server on localhost:8332
	// fmt.Println("Server is running now on http://localhost:8332/")
	// err := http.ListenAndServe(":8332", handlerWithCORS)
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }
}
