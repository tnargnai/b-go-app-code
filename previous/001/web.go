package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    host, _ := os.Hostname()
    fmt.Fprintf(w, "Build 001 - Hi there, I'm served from %s!", host)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8484", nil)
}
