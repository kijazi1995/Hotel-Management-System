package models

type Guest struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IDProof  string `json:"id_proof"`

	Bookings []Booking
	Reviews  []Review
}
