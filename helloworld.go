package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi, RageSquid!")
}

func main() {
    http.HandleFunc("/ragesquid", handler)
    http.ListenAndServe(":8080", nil)
}