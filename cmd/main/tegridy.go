package main

import (
    "log"
    "net/http"
    "github.com/saurabh/tegridyfarms/pkg/controller"
)

func main() {
    http.HandleFunc("/home", controller.HomeHandler)
    http.HandleFunc("/hello", controller.HelloHandler)
    http.HandleFunc("/health", controller.HealthHandler)
    http.HandleFunc("/outdoor", controller.OutdoorHandler)
    http.HandleFunc("/products", controller.GetProducts)
    err := http.ListenAndServe("127.0.0.1:8080", nil)
    if err != nil {
        log.Fatalf("Unable to start HTTP server: %v", err)
    }
}
