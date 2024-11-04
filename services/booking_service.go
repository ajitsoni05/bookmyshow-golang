package services

import (
	"bookmyshow-golang/models"
	"fmt"
)

type BookingService interface {
	CreateBooking(booking models.Booking) (int, error)
	GetBookingByID(id int) (models.Booking, error)
	CancelBooking(id int) error
}

type bookingService struct {
	bookings []models.Booking
}

func NewBookingService() BookingService {
	return &bookingService{
		bookings: []models.Booking{},
	}
}

func (s *bookingService) CreateBooking(booking models.Booking) (int, error) {
	booking.ID = len(s.bookings) + 1
	s.bookings = append(s.bookings, booking)
	return booking.ID, nil
}

func (s *bookingService) GetBookingByID(id int) (models.Booking, error) {
	for _, booking := range s.bookings {
		if booking.ID == id {
			return booking, nil
		}
	}
	return models.Booking{}, fmt.Errorf("booking not found")
}

func (s *bookingService) CancelBooking(id int) error {
	for i, booking := range s.bookings {
		if booking.ID == id {
			s.bookings[i].Status = "canceled"
			return nil
		}
	}
	return fmt.Errorf("booking not found")
}
