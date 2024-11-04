// services/event_service.go
package services

import (
	"bookmyshow-golang/config"
	"bookmyshow-golang/models"
	"database/sql"
	"fmt"
	"log"
)

type EventService interface {
	CreateEvent(event models.Event) (int, error)
	GetEventByID(id int) (models.Event, error)
	GetAllEvents() ([]models.Event, error)
}

type eventService struct {
	db *sql.DB
}

func NewEventService() EventService {
	return &eventService{
		db: config.DB,
	}
}

func (s *eventService) CreateEvent(event models.Event) (int, error) {
	result, err := config.DB.Exec("INSERT INTO events (name, location, datetime, description, category) VALUES (?, ?, ?, ?, ?)",
		event.Name, event.Location, event.DateTime, event.Description, event.Category)
	if err != nil {
		log.Printf("Error creating event: %v", err)
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}

func (s *eventService) GetEventByID(id int) (models.Event, error) {
	var event models.Event
	query := "SELECT id, name, location, datetime, description, category FROM events WHERE id = ?"
	err := s.db.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.Description, &event.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Event{}, fmt.Errorf("event not found")
		}
		return models.Event{}, err
	}
	return event, nil
}

func (s *eventService) GetAllEvents() ([]models.Event, error) {
	query := "SELECT id, name, location, datetime, description, category FROM events"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []models.Event{}
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.Description, &event.Category); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
