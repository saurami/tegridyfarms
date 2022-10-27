package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestContextRoot(t *testing.T) {
    request, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatalf("Unable to reach endpoint: %v", err)
    }

    response := httptest.NewRecorder()

    handler := http.HandlerFunc(contextRoot)
    handler.ServeHTTP(response, request)

    t.Run("Content", func(t *testing.T) {
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
            t.Errorf("Incorrect content type in header ... got %v, want %v", got, want)
        }
    })
}
