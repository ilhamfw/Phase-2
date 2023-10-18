package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
    _ "github.com/go-sql-driver/mysql"
)

type Inventory struct {
    ID          int    `json:"ID"`
    Name        string `json:"Name"`
    ItemCode    string `json:"ItemCode"`
    Stock       int    `json:"Stock"`
    Description string `json:"Description"`
    Status      string `json:"Status"`
}

var db *sql.DB

func main() {
    var err error
    // Membuka koneksi ke database MySQL
    db, err = sql.Open("mysql", "root:@tcp(localhost:3307)/avengers_inventory")
    if err != nil {
        fmt.Println("Failed to connect to the database:", err)
        return
    }
    defer db.Close()

    // Membuat router HTTP
    router := httprouter.New()

    // Mendefinisikan route untuk mengambil daftar inventaris
    router.GET("/inventories", getInventories)

    // Mendefinisikan route untuk mengambil detail inventaris berdasarkan ID
    router.GET("/inventories/:id", getInventoryByID)

    // Mendefinisikan route untuk menambahkan inventaris baru
    router.POST("/inventories", createInventory)

    // Mendefinisikan route untuk memperbarui inventaris berdasarkan ID
    router.PUT("/inventories/:id", updateInventory)

    // Mendefinisikan route untuk menghapus inventaris berdasarkan ID
    router.DELETE("/inventories/:id", deleteInventory)

    port := ":8080"
    fmt.Printf("Server started on port %s\n", port)
    http.ListenAndServe(port, router)
}

// Handler untuk menampilkan daftar inventaris
func getInventories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    rows, err := db.Query("SELECT ID, Name, ItemCode, Stock, Description, Status FROM Inventories")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    inventories := []Inventory{}
    for rows.Next() {
        var inv Inventory
        err := rows.Scan(&inv.ID, &inv.Name, &inv.ItemCode, &inv.Stock, &inv.Description, &inv.Status)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        inventories = append(inventories, inv)
    }

    // Mengembalikan data inventaris sebagai response JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(inventories)
}

// Handler untuk menampilkan detail inventaris berdasarkan ID
func getInventoryByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    // Mendapatkan ID dari parameter
    id := params.ByName("id")

    row := db.QueryRow("SELECT ID, Name, ItemCode, Stock, Description, Status FROM Inventories WHERE ID = ?", id)

    var inv Inventory
    err := row.Scan(&inv.ID, &inv.Name, &inv.ItemCode, &inv.Stock, &inv.Description, &inv.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    // Mengembalikan data inventaris sebagai response JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(inv)
}

// Handler untuk menambahkan inventaris baru
func createInventory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    // Mendapatkan data dari body request
    var inv Inventory
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&inv); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


    // Query database untuk memasukkan inventaris baru
    result, err := db.Exec("INSERT INTO Inventories (Name, ItemCode, Stock, Description, Status) VALUES (?, ?, ?, ?, ?)",
        inv.Name, inv.ItemCode, inv.Stock, inv.Description, inv.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Mengembalikan respons yang memberi tahu bahwa inventaris telah berhasil ditambahkan
    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, "Inventaris berhasil ditambahkan")

    // (Opsional) Mengambil ID dari inventaris yang baru saja ditambahkan
    insertedID, _ := result.LastInsertId()
    inv.ID = int(insertedID)

    // (Opsional) Mengembalikan data inventaris yang baru saja ditambahkan sebagai JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(inv)
}


// Handler untuk memperbarui inventaris berdasarkan ID
func updateInventory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    // Mendapatkan ID dari parameter URL
    id := params.ByName("id")

    // Mendapatkan data inventaris dari body request
    var updatedInv Inventory
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&updatedInv); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Membuka koneksi ke database
    db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/avengers_inventory")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Query database untuk memperbarui inventaris berdasarkan ID
    _, err = db.Exec("UPDATE Inventories SET Name=?, ItemCode=?, Stock=?, Description=?, Status=? WHERE ID=?", updatedInv.Name, updatedInv.ItemCode, updatedInv.Stock, updatedInv.Description, updatedInv.Status, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Mengembalikan respons yang memberi tahu bahwa inventaris berhasil diperbarui
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Inventaris berhasil diperbarui")
}


// Handler untuk menghapus inventaris berdasarkan ID
func deleteInventory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    // Mendapatkan ID dari parameter URL
    id := params.ByName("id")

    // Membuka koneksi ke database
    db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/avengers_inventory")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Query database untuk menghapus inventaris berdasarkan ID
    _, err = db.Exec("DELETE FROM Inventories WHERE ID=?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Mengembalikan respons yang memberi tahu bahwa inventaris berhasil dihapus
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Inventaris berhasil dihapus")
}

