package main

import (
    "log"
    "net/http"
    "html/template"
    "io"
    "os"
    //"image/jpeg"
)

type OutdoorSign struct {
    Title string
    Slogan string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        io.WriteString(w, "Hello, World!")
    }
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/health" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        io.WriteString(w, `{"server": "running", "status": "operational"}`)
    }
}

func outdoorHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/outdoor" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        page := OutdoorSign{Title: "Tegridy Farms", Slogan: "Farming with Tegridy"}
        tmpl := template.Must(template.ParseFiles("outdoor.html"))
        err := tmpl.Execute(w, page)
        if err != nil {
            log.Printf("Unable to parse HTML ... %v", err)
        }
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/home" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        file, err := os.Open("./tegridy_home.jpeg")
        if err != nil {
            log.Printf("Unable to open file: %v", err)
        }
        defer file.Close()
        w.Header().Set("Content-Type", "image/jpeg")
        io.Copy(w, file)
    }
}

func main() {
    http.HandleFunc("/home", homeHandler)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/outdoor", outdoorHandler)
    err := http.ListenAndServe("localhost:8080", nil)
    if err != nil {
        log.Fatalf("Unable to start HTTP server: %v", err)
    }
}
