package main

import (
    "testing"
)

func TestSqliteConnect(t *testing.T) {
    db := sqliteConnect()
    db.SetMaxIdleConns(0)

    t.Run("Ping database", func(t *testing.T) {
        if db.Ping() != nil {
            t.Errorf("Unable to ping in-memory sqlite database ... %v", db.Ping())
        }
    })

    db.Close()

    t.Run("Open Connections", func(t *testing.T) {
        if db.Stats().OpenConnections != 0 {
            t.Errorf("Umable to close database connections ... %v", db.Stats().OpenConnections)
        }
    })
}

func TestReadProducts(t *testing.T) {
    got, err := readProducts()

    t.Run("Available product", func(t *testing.T) {
        var result bool = false
        for _, product := range got {
            if product.Name == "Tegridy Jungle Bud" {
                result = true
                break
            }
        }

        if !result {
            t.Errorf("Product not present in data ... %v", got)
        }
    })

    t.Run("No error", func(t *testing.T) {
        if err != nil {
            t.Errorf("Error is not nil ... %v", err)
        }
    })
}
