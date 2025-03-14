package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

const Port = ":8080"

type Numbers struct {
    A int
    B int
}

func main() {
    mux := http.NewServeMux()

    // Home handler
    mux.HandleFunc("GET /", homeHandler)

    // Decode handler test
    mux.HandleFunc("POST /decode", decodeHandler)

    fmt.Printf("Listening on port %s...\n", Port[1:])
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Path %s : received GET", r.URL.Path)
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
    var n Numbers

    fmt.Printf("%+v\n", r.Body)

    err := json.NewDecoder(r.Body).Decode(&n)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "numbers: %+v\n", n)
}