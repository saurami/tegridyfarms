package controller

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
    "github.com/saurabh/tegridyfarms/pkg/model"
)

func TestGetProducts(t *testing.T) {
    request, err := http.NewRequest("GET", "/products", nil)
    if err != nil {
        t.Fatalf("Unable to reach endpoint: %v", err)
    }

    response := httptest.NewRecorder()

    handler := http.HandlerFunc(GetProducts)
    handler.ServeHTTP(response, request)

    t.Run("Response Code", func(t *testing.T) {
        got := response.Code
        want := http.StatusOK
        if got != want {
            t.Errorf("Incorrect response code ... got %v; want %v", got, want)
        }
    })

    t.Run("Header", func(t *testing.T) {
        got := response.Header().Get("Content-Type")
        want := "application/json"
        if got != want {
            t.Errorf("Incorrect response header ... got %v; want %v", got, want)
        }
    })

    t.Run("Content", func(t *testing.T) {
        err := json.NewDecoder(response.Body).Decode(&model.Products) 
        if err != nil {
            t.Errorf("Error parsing response data into defined product structure: %v", err)
        }
    })
}
