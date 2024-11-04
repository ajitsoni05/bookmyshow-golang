# Bookmyshow-golang
A Go Solution to classic Low Level Design (LLD) problem - BookMyShow


# Folder Structure

bookmyshow/
├── main.go                 # Entry point of the application
├── config/                 # Configuration settings
├── models/                 # Data models for the services
│   ├── event.go            # Model for events (movies, concerts, etc.)
│   ├── booking.go          # Model for user bookings
│   └── user.go             # Model for user information
├── services/               # Service layer implementations
│   ├── event_service.go    # Business logic for managing events
│   ├── booking_service.go  # Logic for handling bookings
│   └── user_service.go     # User management logic
├── handlers/               # API handlers for handling HTTP requests
│   ├── event_handler.go    # Handlers for event-related endpoints
│   ├── booking_handler.go  # Handlers for booking-related endpoints
│   └── user_handler.go     # Handlers for user-related endpoints
├── utils/                  # Utility functions and helpers
└── db.go                   # Database setup and connection logic


