package main

import (
    "log"
    "net/http"
    "io"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    io.WriteString(w, "Welcome to Tegridy Farms!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    io.WriteString(w, `{"server": "running", "status": "operational"}`)
}

func main() {
    http.HandleFunc("/welcome", welcomeHandler)
    http.HandleFunc("/health", healthHandler)
    err := http.ListenAndServe("localhost:8080", nil)
    if err != nil {
        log.Fatalf("Unable to start HTTP server: %v", err)
    }
}
