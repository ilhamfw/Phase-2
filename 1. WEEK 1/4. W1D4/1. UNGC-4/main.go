package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

// Model untuk kejadian kriminal
type CriminalReport struct {
	ID          int    `json:"id"`
	HeroID      int    `json:"hero_id"`
	VillainID   int    `json:"villain_id"`
	Description string `json:"description"`
	EventTime   string `json:"event_time"`
}

var db *sql.DB

func main() {
	var err error
	// Inisialisasi koneksi database
	db, err = sql.Open("mysql", "root:@tcp(localhost:3307)/avengercorp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Post("/criminal-reports", createCriminalReport)
	r.Get("/criminal-reports/{id}", getCriminalReport)
	r.Put("/criminal-reports/{id}", updateCriminalReport)
	r.Delete("/criminal-reports/{id}", deleteCriminalReport)

	// Menambahkan pesan server sedang berjalan
	fmt.Println("Server started on port :8080")

	http.ListenAndServe(":8080", r)
}

func createCriminalReport(w http.ResponseWriter, r *http.Request) {
	// Ambil data dari body request
	var report CriminalReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi data (tambahkan validasi sesuai kebutuhan Anda)

	// Simpan data ke database
	_, err := db.Exec("INSERT INTO CriminalReports (HeroID, VillainID, Description, EventTime) VALUES (?, ?, ?, ?)",
		report.HeroID, report.VillainID, report.Description, report.EventTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons berhasil
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Criminal Report berhasil dibuat")
}

func getCriminalReport(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari parameter URL
	reportID := chi.URLParam(r, "id")

	// Cari data kejadian kriminal berdasarkan ID
	var report CriminalReport
	err := db.QueryRow("SELECT ID, HeroID, VillainID, Description, EventTime FROM CriminalReports WHERE ID = ?", reportID).
		Scan(&report.ID, &report.HeroID, &report.VillainID, &report.Description, &report.EventTime)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Criminal Report not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Kirim data kejadian kriminal sebagai respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func updateCriminalReport(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari parameter URL
	reportID := chi.URLParam(r, "id")

	// Ambil data dari body request
	var updatedReport CriminalReport
	if err := json.NewDecoder(r.Body).Decode(&updatedReport); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi data (tambahkan validasi sesuai kebutuhan Anda)

	// Perbarui data kejadian kriminal berdasarkan ID
	_, err := db.Exec("UPDATE CriminalReports SET HeroID=?, VillainID=?, Description=?, EventTime=? WHERE ID=?", updatedReport.HeroID, updatedReport.VillainID, updatedReport.Description, updatedReport.EventTime, reportID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons berhasil
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Criminal Report berhasil diperbarui")
}

func deleteCriminalReport(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari parameter URL
	reportID := chi.URLParam(r, "id")

	// Hapus data kejadian kriminal berdasarkan ID
	_, err := db.Exec("DELETE FROM CriminalReports WHERE ID=?", reportID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons berhasil
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Criminal Report berhasil dihapus")
}

