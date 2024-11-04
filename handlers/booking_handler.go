package handlers

import (
	"bookmyshow-golang/models"
	"bookmyshow-golang/services"
	"encoding/json"
	"net/http"
)

var bookingService = services.NewBookingService()

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var booking models.Booking
		json.NewDecoder(r.Body).Decode(&booking)
		id, _ := bookingService.CreateBooking(booking)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})
	}
}
