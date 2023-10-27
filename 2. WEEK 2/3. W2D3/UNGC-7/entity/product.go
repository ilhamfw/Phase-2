package entity

import "github.com/jinzhu/gorm"

type Product struct {
    gorm.Model
    Name        string `gorm:"not null"`
    Description string `gorm:"type:varchar(255);not null"`
    ImageURL    string `gorm:"type:varchar(255);not null"`
    Price       int    `gorm:"not null;check:price >= 1000"`
    StoreID     uint   `gorm:"not null"`
}
