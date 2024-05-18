package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "poc-go-dynamodb/models"
    "poc-go-dynamodb/utils"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    item, err := utils.GetItem(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(item)
}
