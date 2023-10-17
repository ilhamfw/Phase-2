package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

// Struct untuk merepresentasikan hero atau villain
type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"imageURL"`
}

func main() {
	// Inisialisasi database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/avengercorp")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Routing endpoint /heroes
	http.HandleFunc("/heroes", func(w http.ResponseWriter, r *http.Request) {
		// Query database untuk mendapatkan data heroes
		rows, err := db.Query("SELECT ID, Name, Universe, Skill, ImageURL FROM Heroes")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		heroes := []Character{}
		for rows.Next() {
			var hero Character
			err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			heroes = append(heroes, hero)
		}

		// Mengembalikan data heroes sebagai response JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(heroes)
	})

	// Routing endpoint /villains
	http.HandleFunc("/villains", func(w http.ResponseWriter, r *http.Request) {
		// Query database untuk mendapatkan data villains
		rows, err := db.Query("SELECT ID, Name, Universe, ImageURL FROM Villain")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		villains := []Character{}
		for rows.Next() {
			var villain Character
			err := rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			villains = append(villains, villain)
		}

		// Mengembalikan data villains sebagai response JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(villains)
	})

	// Menjalankan server pada port tertentu
	port := ":8080"
	fmt.Printf("Server started on port %s\n", port)
	http.ListenAndServe(port, nil)
}
