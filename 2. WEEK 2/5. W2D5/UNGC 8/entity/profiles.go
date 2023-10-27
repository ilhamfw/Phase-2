package entity

import (
	"time"
)

type Profile struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	FirstName  string    `gorm:"type:varchar(50)" json:"first_name"`
	LastName   string    `gorm:"type:varchar(50)" json:"last_name"`
	Address    string    `gorm:"type:varchar(255)" json:"address"`
	PhoneNumber string   `gorm:"type:varchar(15)" json:"phone_number"`
	UserID     uint      `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}