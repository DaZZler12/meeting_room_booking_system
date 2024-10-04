package models

import "sync"

// Room represents a meeting room in the system.
type Room struct {
	RoomID   string
	Name     string
	Capacity int
	mutex    sync.Mutex // To handle concurrent bookings
}

// Lock locks the room to prevent concurrent access.
func (r *Room) Lock() {
	r.mutex.Lock()
}

// Unlock unlocks the room after booking is done.
func (r *Room) Unlock() {
	r.mutex.Unlock()
}
