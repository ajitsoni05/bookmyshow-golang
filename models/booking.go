// models/booking.go
package models

import "time"

type Booking struct {
	ID          int
	EventID     int
	UserName    string
	Seats       int
	BookingTime time.Time
}
