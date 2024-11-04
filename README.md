# Bookmyshow-golang
A Go Solution to classic Low Level Design (LLD) problem - BookMyShow

/
Folder Structure

bookmyshow/
├── main.go                 # Entry point
├── config/                 # Configuration settings
├── models/                 # Data models for the services
│   ├── event.go
│   ├── booking.go
│   └── user.go
├── services/               # Service implementations
│   ├── event_service.go
│   ├── booking_service.go
│   └── user_service.go
├── handlers/               # API handlers for HTTP requests
│   ├── event_handler.go
│   ├── booking_handler.go
│   └── user_handler.go
└── utils/                  # Utilities and helpers
└── db.go               # Database setup and connection

/