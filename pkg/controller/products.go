package controller

import (
    "net/http"
    "encoding/json"
    "strings"
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

func RetrieveProduct(w http.ResponseWriter, r *http.Request) {
    productName := strings.TrimPrefix(r.URL.Path, "/products/")
    if productName == "" {
        http.Error(w, "Invalid product name", http.StatusNotFound)
    } else if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    } else {
        w.Header().Set("Content-Type", "application/json")
        for _, product := range model.Products {
            if product.Name == productName {
                json.NewEncoder(w).Encode(product)
            }
        }
    }
}
