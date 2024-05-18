package handlers

import (
    "encoding/json"
    "net/http"

    "poc-go-dynamodb/models"
    "poc-go-dynamodb/utils"
)

func PutItem(w http.ResponseWriter, r *http.Request) {
    var item models.Item
    err := json.NewDecoder(r.Body).Decode(&item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = utils.PutItem(item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
