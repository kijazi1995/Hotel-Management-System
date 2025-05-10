package models

import "time"

type Booking struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	GuestID  uint      `json:"guest_id"`
	RoomID   uint      `json:"room_id"`
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
	Status   string    `json:"status"` // confirmed, checked_in, checked_out, cancelled

	Guest    Guest
	Room     Room
	Services []RoomService
	//Payment  Payment
}
