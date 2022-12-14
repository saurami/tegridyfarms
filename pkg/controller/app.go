package controller

import (
    "log"
    "net/http"
    "html/template"
    "io"
    "os"
    "github.com/saurabh/tegridyfarms/pkg/model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/home" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        file, err := os.Open("../../web/static/tegridy_home.jpeg")
        if err != nil {
            log.Printf("Unable to open file: %v", err)
        }
        defer file.Close()
        w.Header().Set("Content-Type", "image/jpeg")
        io.Copy(w, file)
    }
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        io.WriteString(w, "Hello, World!")
    }
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/health" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        io.WriteString(w, `{"server": "running", "status": "operational"}`)
    }
}

func OutdoorHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/outdoor" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        page := model.OutdoorSign{Title: "Tegridy Farms", Slogan: "Farming with Tegridy"}
        tmpl := template.Must(template.ParseFiles("../../web/template/outdoor.html"))
        err := tmpl.Execute(w, page)
        if err != nil {
            log.Printf("Unable to parse HTML ... %v", err)
        }
    }
}

