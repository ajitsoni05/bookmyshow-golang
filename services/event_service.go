package services

import (
	"bookmyshow-golang/models"
	"fmt"
)

type EventService interface {
	CreateEvent(event models.Event) (int, error)
	GetEventByID(id int) (models.Event, error)
	GetAllEvents() ([]models.Event, error)
}
type eventService struct {
	events []models.Event
}

// Method to register a new EventService
func NewEventService() EventService {
	return &eventService{
		events: []models.Event{},
	}
}

func (s *eventService) CreateEvent(event models.Event) (int, error) {
	event.ID = len(s.events) + 1
	s.events = append(s.events, event)
	return event.ID, nil
}
func (s *eventService) GetEventByID(id int) (models.Event, error) {
	for _, event := range s.events {
		if event.ID == id {
			return event, nil
		}
	}
	return models.Event{}, fmt.Errorf("Event not found")
}

func (s *eventService) GetAllEvents() ([]models.Event, error) {
	return s.events, nil
}
