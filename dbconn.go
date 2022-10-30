package main

import (
    "log"
    "database/sql"
    _ "modernc.org/sqlite"
)

func sqliteConnect() (sqliteDB *sql.DB) {
    sqliteDB, err := sql.Open("sqlite", ":memory:")
    if err != nil {
        log.Fatalf("Cannot connect to sqlite DB: %v", err)
    }

    return sqliteDB
}
