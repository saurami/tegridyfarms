package routes

import (
	"github.com/saurabh/tegridyfarms/pkg/controller"
	"net/http"
)

func RegisteredPaths(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controller.RetrieveProduct(w, r)
	} else if r.Method == "PUT" {
		controller.UpdateProduct(w, r)
	} else if r.Method == "DELETE" {
		controller.DeleteProduct(w, r)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
