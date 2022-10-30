package main

import "testing"

func TestSqliteConnect(t *testing.T) {
    db := sqliteConnect()
    db.SetMaxIdleConns(0)

    t.Run("Ping database", func(t *testing.T) {
        if db.Ping() != nil {
            t.Errorf("Unable to ping database ...")
        }
    })

    db.Close()
    
    t.Run("Open Connections", func(t *testing.T) {
        got := db.Stats().OpenConnections
        if got != 0 {
            t.Errorf("Umable to close database connections ... %v", db.Stats().OpenConnections)
        }
    })
}
