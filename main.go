package main

import (
	"bookmyshow-golang/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/events", handlers.EventHandler)
	http.HandleFunc("/bookings", handlers.BookingHandler)

	// Start the server
	log.Println("Starting server on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
