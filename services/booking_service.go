// services/booking_service.go
package services

import (
	"bookmyshow-golang/config"
	"bookmyshow-golang/models"
	"database/sql"
	"fmt"
)

type BookingService interface {
	CreateBooking(booking models.Booking) (int, error)
	GetBookingByID(id int) (models.Booking, error)
	GetBookingsByEventID(eventID int) ([]models.Booking, error)
}

type bookingService struct {
	db *sql.DB
}

func NewBookingService() BookingService {
	return &bookingService{
		db: config.DB,
	}
}

func (s *bookingService) CreateBooking(booking models.Booking) (int, error) {
	query := "INSERT INTO bookings (event_id, user_name, seats) VALUES (?, ?, ?)"
	result, err := s.db.Exec(query, booking.EventID, booking.UserName, booking.Seats)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *bookingService) GetBookingByID(id int) (models.Booking, error) {
	var booking models.Booking
	query := "SELECT id, event_id, user_name, seats, booking_time FROM bookings WHERE id = ?"
	err := s.db.QueryRow(query, id).Scan(&booking.ID, &booking.EventID, &booking.UserName, &booking.Seats, &booking.BookingTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Booking{}, fmt.Errorf("booking not found")
		}
		return models.Booking{}, err
	}
	return booking, nil
}

func (s *bookingService) GetBookingsByEventID(eventID int) ([]models.Booking, error) {
	query := "SELECT id, event_id, user_name, seats, booking_time FROM bookings WHERE event_id = ?"
	rows, err := s.db.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := []models.Booking{}
	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.ID, &booking.EventID, &booking.UserName, &booking.Seats, &booking.BookingTime); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bookings, nil
}
