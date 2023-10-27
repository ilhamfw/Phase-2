// swag init --parseDependency --parseInternal
package config

import (
    "fmt"
    "github.com/joho/godotenv"
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // Import driver PostgreSQL
)

func InitDB() (*gorm.DB, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, err
    }

    dbDriver := "postgres" // Menggunakan driver PostgreSQL
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    
    dbName := os.Getenv("DB_NAME")

    dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

    db, err := gorm.Open(dbDriver, dbPath)
    if err != nil {
        return nil, err
    }
    return db, nil
}
