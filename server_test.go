package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    request, err := http.NewRequest("GET", "/hello", nil)
    if err != nil {
        t.Fatalf("Unable to reach endpoint: %v", err)
    }

    response := httptest.NewRecorder()

    handler := http.HandlerFunc(helloHandler)
    handler.ServeHTTP(response, request)

    t.Run("Content", func(t *testing.T){
        got := response.Body.String()
        want := "Hello, World!"
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

func TestOutdoorHandler(t *testing.T) {
    request, err := http.NewRequest("GET", "/outdoor", nil)
    if err != nil {
        t.Fatalf("Unable to reach endpoint: %v", err)
    }

    response := httptest.NewRecorder()

    handler := http.HandlerFunc(outdoorHandler)
    handler.ServeHTTP(response, request)

    t.Run("Response Code", func(t *testing.T) {
        got := response.Code
        want := http.StatusOK
        if got != want {
            t.Errorf("Unexpected response code ... got %v, want %v", got, want)
        }
    })

    t.Run("Header", func(t *testing.T) {
        got := response.Header().Get("Content-Type")
        want := "text/html; charset=utf-8"
        if got != want {
            t.Errorf("Incorrect response header ... got %v, want %v", got, want)
        }
    })

    t.Run("Content", func(t *testing.T) {
        got := response.Body.String()
        if !strings.Contains(got, "Tegridy Farms") {
            t.Errorf("Outdoor sign for 'Tegridy Farms' missing... got %v", got)
        }
    })
}
