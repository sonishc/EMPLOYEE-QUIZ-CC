package controllers

import (
	"../db"
	"../utils"
	"encoding/json"
	"net/http"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	json.NewEncoder(w).Encode(db.GetQuestions())
}
