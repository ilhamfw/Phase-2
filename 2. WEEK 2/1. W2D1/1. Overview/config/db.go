package config

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root@tcp(localhost:3307)/overview_branch")
    if err != nil {
        return nil, err
    }
    return db, nil
}
