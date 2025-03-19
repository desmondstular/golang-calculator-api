package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

const Port = ":8080"

type Numbers struct {
    A float32
    B float32
}

func main() {
    mux := http.NewServeMux()

    // Home handler
    mux.HandleFunc("GET /", homeHandler)

    // Operation routes
    mux.HandleFunc("POST /add", addHandler)
    mux.HandleFunc("POST /subtract", subtractHandler)
    mux.HandleFunc("POST /multiply", multiplyHandler)
    mux.HandleFunc("POST /divide", divideHandler)

    // Decode handler test
    mux.HandleFunc("POST /decode", decodeHandler)

    fmt.Printf("Listening on port %s...\n", Port[1:])
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Path %s : received GET", r.URL.Path)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
    var n Numbers
    var value float32

    err := json.NewDecoder(r.Body).Decode(&n)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    value = n.A + n.B

    fmt.Fprintf(w, "%v\n", value)
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
    var n Numbers
    var value float32

    err := json.NewDecoder(r.Body).Decode(&n)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    value = n.A - n.B

    fmt.Fprintf(w, "%v\n", value)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
    var n Numbers
    var value float32

    err := json.NewDecoder(r.Body).Decode(&n)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    value = n.A * n.B

    fmt.Fprintf(w, "%v\n", value)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
    var n Numbers
    var value float32

    err := json.NewDecoder(r.Body).Decode(&n)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Cannot divide by zero
    if n.B == 0 {
        http.Error(w, "Unable to divide by zero", http.StatusBadRequest)
        return
    }

    value = n.A / n.B

    fmt.Fprintf(w, "%v\n", value)
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