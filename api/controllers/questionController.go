package controllers

import (
    "../models"
    "../utils"
    "encoding/json"
    "net/http"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
    utils.SetHeaders(w)
    json.NewEncoder(w).Encode(models.GetQuestions())
}
