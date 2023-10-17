package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Buat koneksi ke database MySQL
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/database_name")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Query untuk mengambil data dari tabel Heroes
    rows, err := db.Query("SELECT ID, Name, Universe, Skill, ImageURL FROM Heroes")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    // Baca data dari hasil query
    for rows.Next() {
        var ID int
        var Name string
        var Universe string
        var Skill string
        var ImageURL string
        if err := rows.Scan(&ID, &Name, &Universe, &Skill, &ImageURL); err != nil {
            panic(err.Error())
        }
        fmt.Printf("ID: %d, Name: %s, Universe: %s, Skill: %s, ImageURL: %s\n", ID, Name, Universe, Skill, ImageURL)
    }
}
