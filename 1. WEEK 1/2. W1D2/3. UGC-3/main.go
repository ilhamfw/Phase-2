package main

import "github.com/julienschmidt/httprouter"


// Handler untuk menampilkan daftar inventaris
func getInventories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    // Membuka koneksi ke database
    db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/avengersinventory")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Query database untuk mendapatkan data inventaris
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

    // Membuka koneksi ke database
    db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/avengersinventory")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Query database untuk mendapatkan data inventaris berdasarkan ID
    rows, err := db.Query("SELECT ID, Name, ItemCode, Stock, Description, Status FROM Inventories WHERE ID = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Membaca data inventaris
    if rows.Next() {
        var inv Inventory
        err := rows.Scan(&inv.ID, &inv.Name, &inv.ItemCode, &inv.Stock, &inv.Description, &inv.Status)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Mengembalikan data inventaris sebagai response JSON
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(inv)
    } else {
        // Mengembalikan response bahwa inventaris tidak ditemukan
        http.Error(w, "Inventory not found", http.StatusNotFound)
    }
}

// Handler untuk menambahkan inventaris baru
func createInventory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    // ... Implementasi handler POST /inventories ...
}

// Handler untuk memperbarui inventaris berdasarkan ID
func updateInventory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    // ... Implementasi handler PUT /inventories/:id ...
}

// Handler untuk menghapus inventaris berdasarkan ID
func deleteInventory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    // ... Implementasi handler DELETE /inventories/:id ...
}


// Contoh operasi database untuk menampilkan daftar inventaris
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

port := ":8080"
fmt.Printf("Server started on port %s\n", port)
http.ListenAndServe(port, router)
