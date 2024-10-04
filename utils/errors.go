package utils

import "errors"

var (
	ErrRoomAlreadyBooked = errors.New("room is already booked during the requested time")
	ErrRoomNotFound      = errors.New("room not found")
	ErrBookingNotFound   = errors.New("booking not found")
)
