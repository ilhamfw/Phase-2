package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"nebula_app/entity"
)

func main() {
	// koneksi database PostgreSQL 
	db, err := gorm.Open("postgres", "host=localhost dbname=nebula_app user=postgres password=postgres sslmode=disable")
	


	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	// Auto Migrate tabel users dan profiles
	db.AutoMigrate(&entity.User{}, &entity.Profile{})

	fmt.Println("Auto migration completed.")
}
