package main

import (
    "fmt"
    "net/http"
)

func contextRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", contextRoot)
    http.ListenAndServe(":8080", nil)
}
