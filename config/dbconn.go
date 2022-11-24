package main

import (
    "log"
    "database/sql"
    _ "modernc.org/sqlite"
)

type Products struct {
    Name string
    Updated string
}

func sqliteConnect() (sqliteDB *sql.DB) {
    // driver and data source names
    sqliteDB, err := sql.Open("sqlite", "./tegridyfarms.db")
    if err != nil {
        log.Fatalf("Cannot connect to sqlite DB: %v", err)
    }

    return sqliteDB
}

func readProducts() ([]Products, error) {
    db := sqliteConnect()
    rows, err := db.Query("SELECT * FROM product")
    if err != nil {
        log.Fatalf("Error retrieving records from the table: %v", err)
    }
    defer rows.Close()

    var product []Products
    for rows.Next() {
        var item Products
        if err := rows.Scan(&item.Name, &item.Updated); err != nil {
            log.Fatalf("Data error on table row: %v", err)
            return nil, err
        }
        product = append(product, item)
    }
    return product, nil
}
