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

    var allProducts []Products
    for rows.Next() {
        var item Products
        if err := rows.Scan(&item.Name, &item.Updated); err != nil {
            log.Fatalf("Data error on table row: %v", err)
            return nil, err
        }
        allProducts = append(allProducts, item)
    }
    return allProducts, nil
}

func retrieveProduct(name string) error {
    db := sqliteConnect()
    row := db.QueryRow("SELECT * FROM product WHERE name=?", name)
    if row == nil {
        log.Fatalf("Error with record: %v", row)
    }

    var product Products
    if err := row.Scan(&product.Name, &product.Updated); err != nil {
        log.Fatalf("Data error in table row: %v", err)
        return err
    }
    return nil
}
