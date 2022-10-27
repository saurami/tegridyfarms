package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestContextRoot(t *testing.T) {
    //request := httptest.NewRequest("GET", "/", nil)
    request, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Errorf("Unable to reach endpoint: %v", err)
    }
    response := httptest.NewRecorder()

    handler := http.HandlerFunc(contextRoot)
    handler.ServeHTTP(response, request)

    if response.Code != http.StatusOK {
        t.Errorf("Unexpected response code... got %v, want 200", response.Code)
    }

    if response.Body.String() != "Hello, World!" {
        t.Errorf("Unexpected response body... got %v, want 'Hello, World!'", response.Body.String())
    }
}
