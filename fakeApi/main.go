package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/products", productsHandler)
    http.HandleFunc("/categories", categoriesHandler)

    http.ListenAndServe(":8080", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(categories)
}
