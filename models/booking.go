package models

import "time"

// Booking represents a room booking by a user.
type Booking struct {
	BookingID string
	Room      *Room
	User      *User
	StartTime time.Time
	EndTime   time.Time
	Status    string
}
