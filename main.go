package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "poc-go-dynamodb/handlers"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/item/{id}", handlers.GetItem).Methods("GET")
    r.HandleFunc("/item", handlers.PutItem).Methods("POST")

    log.Println("Servidor iniciado na porta :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
