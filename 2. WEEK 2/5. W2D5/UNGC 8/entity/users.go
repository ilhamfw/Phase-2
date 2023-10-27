package entity

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Profiles  []Profile `json:"profiles"`
}