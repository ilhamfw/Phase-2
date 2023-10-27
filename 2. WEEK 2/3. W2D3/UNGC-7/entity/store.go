package entity

import "github.com/jinzhu/gorm"

type Store struct {
    gorm.Model
    StoreEmail string `gorm:"type:varchar(100);unique_index;not null"`
    Password   string `gorm:"type:varchar(100);not null"`
    StoreName  string `gorm:"type:varchar(100);not null"`
    StoreType  string `gorm:"type:varchar(10);default:'silver'"`
}

type LoginData struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}
