package models

import "time"

type RoomService struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BookingID   uint      `json:"booking_id"`
	ServiceName string    `json:"service_name"`
	Cost        float64   `json:"cost"`
	ServiceDate time.Time `json:"service_date"`

	Booking Booking
}
