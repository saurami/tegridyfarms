package controller

import (
	"encoding/json"
	"github.com/saurabh/tegridyfarms/pkg/model"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
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
	} else {
		w.Header().Set("Content-Type", "application/json")
		for _, product := range model.Products {
			if product.Name == productName {
				json.NewEncoder(w).Encode(product)
			}
		}
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/product" {
		http.Error(w, "Invalid path", http.StatusNotFound)
	} else if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	} else {
		w.Header().Set("Content-Type", "application/json")
		var newProduct model.Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "Unable to parse given data to specified data format", http.StatusNotAcceptable)
		} else {
			model.Products = append(model.Products, newProduct)
			json.NewEncoder(w).Encode(newProduct)
		}
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productName := strings.TrimPrefix(r.URL.Path, "/products/")
	if productName == "" {
		http.Error(w, "Invalid product name", http.StatusNotFound)
	} else {
		//    TO DO
		w.Header().Set("Content-Type", "application/json")
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productName := strings.TrimPrefix(r.URL.Path, "/products/")
	if productName == "" {
		http.Error(w, "Invalid product name", http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		for _, product := range model.Products {
			if product.Name == productName {
				// get index of product name
				// use index to call delete
				//i := slices.Index(model.Products, product.Name)
				i := slices.IndexFunc(model.Products, func(p model.Product) bool {
					return p.Name == product.Name
				})
				model.Products = slices.Delete(model.Products, i, i+1)
				break
			}
		}
		json.NewEncoder(w).Encode(model.Products)
	}
}
