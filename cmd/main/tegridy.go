package main

import (
	"github.com/saurabh/tegridyfarms/pkg/controller"
	"github.com/saurabh/tegridyfarms/pkg/routes"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", controller.HomeHandler)
	http.HandleFunc("/hello", controller.HelloHandler)
	http.HandleFunc("/health", controller.HealthHandler)
	http.HandleFunc("/outdoor", controller.OutdoorHandler)
	http.HandleFunc("/products", controller.GetProducts)
	http.HandleFunc("/products/", routes.RegisteredPaths)
	http.HandleFunc("/product", controller.CreateProduct)

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalf("Unable to start HTTP server: %v", err)
	}
}
