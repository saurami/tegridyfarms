package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Write([]byte(`{"server": "running", "status": "operational"}`))
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/health", healthHandler)
    err := http.ListenAndServe("localhost:8080", nil)
    if err != nil {
        log.Fatalf("Unable to start HTTP server: %v", err)
    }
}
