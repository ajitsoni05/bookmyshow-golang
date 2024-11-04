package models

type Event struct {
	ID          int
	Name        string
	Location    string
	DateTime    string
	Description string
	Category    string
}

type Seat struct {
	ID       int
	EventID  int
	Row      int
	Number   int
	IsBooked bool
}
