package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	json.NewEncoder(w).Encode(models.GetQuestions())
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(models.GetQuestion(params["id"]))
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	var question models.Question
	_ = json.NewDecoder(r.Body).Decode(&question)
	message, status := question.Validate()
	if status {
		models.CreateQuestion(&question)
	}
	json.NewEncoder(w).Encode(message)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	params := mux.Vars(r)
	question := models.GetQuestion(params["id"])
	_ = json.NewDecoder(r.Body).Decode(&question)
	message, status := question.Validate()
	if status {
		models.UpdateQuestion(question)
	}
	json.NewEncoder(w).Encode(message)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	params := mux.Vars(r)
	question := models.GetQuestion(params["id"])
	models.DeleteQuestion(question)
	json.NewEncoder(w).Encode("success")
}
