package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestWelcomeHandler(t *testing.T) {
    request, err := http.NewRequest("GET", "/welcome", nil)
    if err != nil {
        t.Fatalf("Unable to reach endpoint: %v", err)
    }

    response := httptest.NewRecorder()

    handler := http.HandlerFunc(welcomeHandler)
    handler.ServeHTTP(response, request)

    t.Run("Content", func(t *testing.T){
        got := response.Body.String()
        want := "Welcome to Tegridy Farms!"
        if got != want {
            t.Errorf("Incorrect content ... got %v, want %v", got, want)
        }
    })

    t.Run("Response Code", func(t *testing.T) {
        got := response.Code
        want := http.StatusOK
        if got != want {
            t.Errorf("Unexpected response code ... got %v, want %v", got, want)
        }
    })

    t.Run("Header", func(t *testing.T) {
        got := response.Header().Get("Content-Type")
        want := "text/plain; charset=utf-8"
        if got != want {
            t.Errorf("Incorrect content header ... got %v, want %v", got, want)
        }
    })
}


func TestHealthHandler(t *testing.T) {
    request, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatalf("Unable to reach endpoint: %v", err)
    }

    response := httptest.NewRecorder()

    handler := http.HandlerFunc(healthHandler)
    handler.ServeHTTP(response, request)

    t.Run("Content", func(t *testing.T){
        got := response.Body.String()
        want := `{"server": "running", "status": "operational"}`
        if got != want {
            t.Errorf("Incorrect content ... got %v, want %v", got, want)
        }
    })

    t.Run("Response Code", func(t *testing.T) {
        got := response.Code
        want := http.StatusOK
        if got != want {
            t.Errorf("Unexpected response code ... got %v, want %v", got, want)
        }
    })

    t.Run("Header", func(t *testing.T) {
        got := response.Header().Get("Content-Type")
        want := "application/json; charset=utf-8"
        if got != want {
            t.Errorf("Incorrect content header ... got %v, want %v", got, want)
        }
    })
}
