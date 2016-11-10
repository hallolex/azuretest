package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi, RageSquid!")
}

func PrintLine() {
	fmt.Println("hello thur")
}

func main() {
	PrintLine()
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}