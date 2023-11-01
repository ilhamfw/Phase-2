package entity

import (
    "time"
    "github.com/go-playground/validator/v10"
)

type User struct {
    ID        uint `gorm:"primaryKey"`
    Username  string `validate:"required,unique"`
    Email     string `validate:"required,email,unique"`
    Password  string `validate:"required,min=6"`
    Age       int `validate:"required,gte=9"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Photo struct {
    ID        uint `gorm:"primaryKey"`
    Title     string `validate:"required"`
    Caption   string
    PhotoURL  string `validate:"required"`
    UserID    uint
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Comment struct {
    ID        uint `gorm:"primaryKey"`
    UserID    uint
    PhotoID   uint
    Message   string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type SocialMedia struct {
    ID            uint `gorm:"primaryKey"`
    Name          string `validate:"required"`
    SocialMediaURL string `validate:"required"`
    UserID        uint
}

var validate *validator.Validate

func init() {
    validate = validator.New()
}
 