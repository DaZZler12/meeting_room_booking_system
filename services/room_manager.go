// Singleton Room Manager

package services

import (
	"errors"
	"sync"

	"meeting_room_booking_system/models"
	"meeting_room_booking_system/utils"
)

// RoomManager is responsible for managing rooms.
type RoomManager struct {
	rooms map[string]*models.Room
	mutex sync.Mutex
}

var roomManagerInstance *RoomManager
var roomManagerOnce sync.Once

// GetRoomManager provides a singleton instance of RoomManager.
func GetRoomManager() *RoomManager {
	roomManagerOnce.Do(func() {
		roomManagerInstance = &RoomManager{
			rooms: make(map[string]*models.Room),
		}
	})
	return roomManagerInstance
}

// AddRoom adds a new room to the system.
func (rm *RoomManager) AddRoom(room *models.Room) error {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	if _, exists := rm.rooms[room.RoomID]; exists {
		return errors.New("room already exists")
	}

	rm.rooms[room.RoomID] = room
	return nil
}

// GetRoom fetches a room by its ID.
func (rm *RoomManager) GetRoom(roomID string) (*models.Room, error) {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	room, exists := rm.rooms[roomID]
	if !exists {
		return nil, utils.ErrRoomNotFound
	}

	return room, nil
}
