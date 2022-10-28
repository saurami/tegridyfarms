package main

import (
    "log"
    "net/http"
    "html/template"
    "io"
)

type OutdoorSign struct {
    Title string
    Slogan string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    io.WriteString(w, "Hello, World!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    io.WriteString(w, `{"server": "running", "status": "operational"}`)
}

func outdoorHandler(w http.ResponseWriter, r *http.Request) {
    page := OutdoorSign{Title: "Tegridy Farms", Slogan: "Farming with Tegridy"}
    tmpl := template.Must(template.ParseFiles("outdoor.html"))
    err := tmpl.Execute(w, page)
    if err != nil {
        log.Printf("Unable to parse HTML ... %v", err)
    }
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/outdoor", outdoorHandler)
    err := http.ListenAndServe("localhost:8080", nil)
    if err != nil {
        log.Fatalf("Unable to start HTTP server: %v", err)
    }
}
