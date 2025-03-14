package main

import (
    "fmt"
    "log"
    "net/http"
)

const Port = ":8080"

func main() {
    mux := http.NewServeMux()

    // Home handler
    mux.HandleFunc("GET /", homeHandler)

    fmt.Printf("Listening on port %s...\n", Port[1:])
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Path %s : received GET", r.URL.Path)
}