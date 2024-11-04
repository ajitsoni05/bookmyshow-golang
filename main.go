// main.go
package main

import (
	"bookmyshow-golang/config"
	"bookmyshow-golang/handlers"
	"log"
	"net/http"
)

func main() {
	config.InitDB() // Initialize the database connection
	defer config.DB.Close()

	// Initialize services and handlers
	http.HandleFunc("/events", handlers.EventHandler)
	http.HandleFunc("/bookings", handlers.BookingHandler)

	// Start server
	log.Println("Server started on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
