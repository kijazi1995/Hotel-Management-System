package models

type Room struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	HotelID    uint    `json:"hotel_id"`
	RoomNumber string  `json:"room_number"`
	Type       string  `json:"type"`
	Price      float64 `json:"price"`
	Status     string  `json:"status"` // available, booked, maintenance

	Hotel     Hotel
	Amenities []Amenity `gorm:"many2many:room_amenities;"`
}
