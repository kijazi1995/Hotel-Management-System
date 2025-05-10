package models

type Amenity struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Rooms []Room `gorm:"many2many:room_amenities;"`
}
