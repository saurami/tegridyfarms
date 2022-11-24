package controller

import (
    "net/http"
    "encoding/json"
    "github.com/saurabh/tegridyfarms/pkg/model"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/products" {
        http.Error(w, "Path doesn't exist", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
    } else {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(model.Products)
    }
}
