package handlers

import (
	"bookmyshow-golang/models"
	"bookmyshow-golang/services"
	"encoding/json"
	"net/http"
	"strconv"
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

	case http.MethodDelete:
		// Assuming `id` parameter is passed in query string
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		err := bookingService.CancelBooking(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
