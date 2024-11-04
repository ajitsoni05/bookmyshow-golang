package handlers

import (
	"bookmyshow-golang/models"
	"bookmyshow-golang/services"
	"encoding/json"
	"net/http"
)

var eventService = services.NewEventService()

func EventHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var event models.Event
		json.NewDecoder(r.Body).Decode(&event)
		id, _ := eventService.CreateEvent(event)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"id": id})

	case http.MethodGet:
		events, _ := eventService.GetAllEvents()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(events)
	}
}
