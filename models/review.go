package models

import "time"

type Review struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	GuestID    uint      `json:"guest_id"`
	RoomID     uint      `json:"room_id"`
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	ReviewDate time.Time `json:"review_date"`

	Guest Guest
	Room  Room
}
