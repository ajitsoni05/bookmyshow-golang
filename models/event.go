// models/event.go
package models

import "time"

type Event struct {
	ID          int       `json:"id"`          // ID from database (primary key)
	Name        string    `json:"name"`        // Name of the event
	Location    string    `json:"location"`    // Location of the event
	DateTime    time.Time `json:"datetime"`    // Event date and time (matches MySQL DATETIME)
	Description string    `json:"description"` // Description of the event
	Category    string    `json:"category"`    // Category like Music, Comedy, etc.
}

type Seat struct {
	ID       int
	EventID  int
	Row      int
	Number   int
	IsBooked bool
}
