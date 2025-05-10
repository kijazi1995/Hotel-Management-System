package models

import "time"

type Payment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BookingID   uint      `json:"booking_id"`
	Amount      float64   `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	Method      string    `json:"method"` // cash, card, online
	Status      string    `json:"status"`

	Booking Booking
}
