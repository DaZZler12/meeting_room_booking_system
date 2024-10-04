
package main

import (
	"fmt"
	"sync"
	"time"

	"meeting_room_booking_system/models"
	"meeting_room_booking_system/services"
)

func main() {
	// Initialize the room manager and booking service
	roomManager := services.GetRoomManager()
	bookingService := services.GetBookingService()

	// Create 10 rooms
	for i := 1; i <= 10; i++ {
		room := &models.Room{
			RoomID:   fmt.Sprintf("R%d", i),
			Name:     fmt.Sprintf("Conference Room %d", i),
			Capacity: 10,
		}
		err := roomManager.AddRoom(room)
		if err != nil {
			fmt.Println("Error adding room:", err)
			return
		}
	}

	// Create users
	users := []*models.User{
		{UserID: "U1", Name: "Alice", Email: "alice@example.com"},
		{UserID: "U2", Name: "Bob", Email: "bob@example.com"},
		{UserID: "U3", Name: "Charlie", Email: "charlie@example.com"},
		{UserID: "U4", Name: "Dave", Email: "dave@example.com"},
		{UserID: "U5", Name: "Eve", Email: "eve@example.com"},
		// {UserID: "U51", Name: "Eve", Email: "eve@example.com"},
		// {UserID: "U52", Name: "Eve", Email: "eve@example.com"},
		// {UserID: "U53", Name: "Eve", Email: "eve@example.com"},
		// {UserID: "U54", Name: "Eve", Email: "eve@example.com"},
		// {UserID: "U55", Name: "Eve", Email: "eve@example.com"},
		// {UserID: "U56", Name: "Eve", Email: "eve@example.com"},
		// {UserID: "U57", Name: "Eve", Email: "eve@example.com"},
		// {UserID: "U58", Name: "Eve", Email: "eve@example.com"},
	}

	// Simulate concurrent booking requests
	var wg sync.WaitGroup
	wg.Add(len(users))

	// Simulate booking rooms concurrently
	for i, user := range users {
		go func(u *models.User, roomIndex int) {
			defer wg.Done()
			room, bookingErr := roomManager.GetRoom(fmt.Sprintf("R%d", roomIndex+1)) // Get different room for each user
			if bookingErr != nil {
				fmt.Printf("Error fetching room for %s: %v\n", u.Name, bookingErr)
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
			err := bookingService.CreateBooking(booking)
			if err != nil {
				fmt.Printf("Booking for %s failed: %v\n", u.Name, err)
			} else {
				fmt.Printf("Booking for %s succeeded in %s\n", u.Name, room.Name)
			}
		}(user, i)
	}

	wg.Wait()
}
