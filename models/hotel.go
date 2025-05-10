package models

type Hotel struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`

	Rooms  []Room  `gorm:"foreignKey:HotelID"`
	Staffs []Staff `gorm:"foreignKey:HotelID"`
}
