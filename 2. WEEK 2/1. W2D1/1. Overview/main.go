package main

import (
    "fmt"
    "net/http"
    "overview/config"
    "overview/handler"
    "overview/repository"
    "github.com/gorilla/mux"
)

func main() {
    db, err := config.SetupDB()
    if err != nil {
        fmt.Println("Failed to connect to the database")
        return
    }
    defer db.Close()

    branchRepo := repository.NewBranchRepository(db)
    branchHandler := handler.NewBranchHandler(branchRepo)

    r := mux.NewRouter()

    r.HandleFunc("/branches", branchHandler.GetBranches).Methods("GET")
    r.HandleFunc("/branches/{branch_id}", branchHandler.GetBranchByID).Methods("GET")
    r.HandleFunc("/branches", branchHandler.AddBranch).Methods("POST")
    r.HandleFunc("/branches/{branch_id}", branchHandler.UpdateBranch).Methods("PUT")
    r.HandleFunc("/branches/{branch_id}", branchHandler.DeleteBranch).Methods("DELETE")

    fmt.Println("Server is running on port 8080")
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
