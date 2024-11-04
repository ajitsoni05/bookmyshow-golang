package models

type Booking struct {
	ID        int
	UserID    int
	EventID   int
	SeatID    int
	Status    string // e.g., "confirmed", "pending", "canceled"
	CreatedAt string
}
