package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"meeting_room_booking_system/models"
	"meeting_room_booking_system/services"
)

// TestConcurrentBookings tests the booking system for race conditions
func TestConcurrentBookings(t *testing.T) {
	roomManager := services.GetRoomManager()
	bookingService := services.GetBookingService()

	// Create a room
	room := &models.Room{
		RoomID:   "R1",
		Name:     "Conference Room 1",
		Capacity: 10,
	}
	err := roomManager.AddRoom(room)
	if err != nil {
		t.Fatalf("Failed to add room: %v", err)
	}

	// Create a user
	user := &models.User{
		UserID: "U1",
		Name:   "Alice",
		Email:  "alice@example.com",
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Simulate concurrent booking requests
	go func() {
		defer wg.Done()
		booking := &models.Booking{
			BookingID: "B1",
			Room:      room,
			User:      user,
			StartTime: time.Now().Add(1 * time.Hour),
			EndTime:   time.Now().Add(2 * time.Hour),
			Status:    "Pending",
		}
		err := bookingService.CreateBooking(booking)
		if err != nil {
			t.Log("Booking 1 failed:", err)
		} else {
			t.Log("Booking 1 succeeded")
		}
	}()

	go func() {
		defer wg.Done()
		booking := &models.Booking{
			BookingID: "B2",
			Room:      room,
			User:      user,
			StartTime: time.Now().Add(1 * time.Hour),
			EndTime:   time.Now().Add(2 * time.Hour),
			Status:    "Pending",
		}
		err := bookingService.CreateBooking(booking)
		if err != nil {
			t.Log("Booking 2 failed:", err)
		} else {
			t.Log("Booking 2 succeeded")
		}
	}()

	wg.Wait()
}

// TestMultipleUsers tests booking from multiple users concurrently
func TestMultipleUsers(t *testing.T) {
	roomManager := services.GetRoomManager()
	bookingService := services.GetBookingService()

	// Create multiple rooms
	for i := 1; i <= 5; i++ {
		room := &models.Room{
			RoomID:   fmt.Sprintf("R%d", i),
			Name:     fmt.Sprintf("Conference Room %d", i),
			Capacity: 10,
		}
		err := roomManager.AddRoom(room)
		if err != nil {
			t.Fatalf("Failed to add room %s: %v", room.RoomID, err)
		}
	}

	// Create users
	users := []*models.User{
		{UserID: "U1", Name: "Alice", Email: "alice@example.com"},
		{UserID: "U2", Name: "Bob", Email: "bob@example.com"},
		{UserID: "U3", Name: "Charlie", Email: "charlie@example.com"},
		{UserID: "U4", Name: "Dave", Email: "dave@example.com"},
		{UserID: "U5", Name: "Eve", Email: "eve@example.com"},
	}

	var wg sync.WaitGroup
	wg.Add(len(users))

	for i, user := range users {
		go func(u *models.User, roomIndex int) {
			defer wg.Done()
			roomID := fmt.Sprintf("R%d", roomIndex%5+1) // Cycle through rooms
			room, err := roomManager.GetRoom(roomID)
			if err != nil {
				t.Logf("Error fetching room for %s: %v\n", u.Name, err)
				return
			}

			booking := &models.Booking{
				BookingID: fmt.Sprintf("B_%s", u.UserID),
				Room:      room,
				User:      u,
				StartTime: time.Now().Add(1 * time.Hour),
				EndTime:   time.Now().Add(2 * time.Hour),
				Status:    "Pending",
			}
			err = bookingService.CreateBooking(booking)
			if err != nil {
				t.Logf("Booking for %s failed: %v\n", u.Name, err)
			} else {
				t.Logf("Booking for %s succeeded in %s\n", u.Name, room.Name)
			}
		}(user, i)
	}

	wg.Wait()
}
