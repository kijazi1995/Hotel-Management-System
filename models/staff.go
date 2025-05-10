package models

import "time"

type Staff struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	HotelID  uint      `json:"hotel_id"`
	FullName string    `json:"full_name"`
	Role     string    `json:"role"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	HireDate time.Time `json:"hire_date"`

	Hotel Hotel
}
