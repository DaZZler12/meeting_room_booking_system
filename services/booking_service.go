// Singleton Booking Service

package services

import (
	"fmt"
	"sync"
	"time"

	"meeting_room_booking_system/models"
	"meeting_room_booking_system/utils"
)

// BookingService is responsible for managing bookings.
type BookingService struct {
	bookings map[string]*models.Booking
	mutex    sync.Mutex
}

var bookingServiceInstance *BookingService
var bookingServiceOnce sync.Once

// GetBookingService provides a singleton instance of BookingService.
func GetBookingService() *BookingService {
	bookingServiceOnce.Do(func() {
		bookingServiceInstance = &BookingService{
			bookings: make(map[string]*models.Booking),
		}
	})
	return bookingServiceInstance
}

// CreateBooking tries to book a room for a user.
func (bs *BookingService) CreateBooking(booking *models.Booking) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	// Lock the room to prevent concurrent bookings
	booking.Room.Lock()
	defer booking.Room.Unlock()

	// Check for time conflicts
	for _, b := range bs.bookings {
		if b.Room.RoomID == booking.Room.RoomID && isOverlapping(b, booking) {
			return utils.ErrRoomAlreadyBooked
		}
	}

	// Save the booking
	bookingID := fmt.Sprintf("%s-%s", booking.User.UserID, booking.Room.RoomID)
	bs.bookings[bookingID] = booking
	booking.Status = "Confirmed"
	return nil
}

// RescheduleBooking updates the booking times.
func (bs *BookingService) RescheduleBooking(bookingID string, startTime, endTime time.Time) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	booking, exists := bs.bookings[bookingID]
	if !exists {
		return utils.ErrBookingNotFound
	}

	// Lock the room during reschedule to prevent concurrent changes
	booking.Room.Lock()
	defer booking.Room.Unlock()

	// Update the booking details
	booking.StartTime = startTime
	booking.EndTime = endTime
	return nil
}

// Helper function to check if two bookings overlap
func isOverlapping(b1, b2 *models.Booking) bool {
	return b1.StartTime.Before(b2.EndTime) && b2.StartTime.Before(b1.EndTime)
}
