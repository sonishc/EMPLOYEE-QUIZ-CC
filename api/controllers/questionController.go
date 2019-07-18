package controllers

import (
	"../db"
	"../models"
	"../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	json.NewEncoder(w).Encode(db.GetQuestions())
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(db.GetQuestion(params["id"]))
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	var question models.Question
	_ = json.NewDecoder(r.Body).Decode(&question)
	message, status := question.Validate()
	if status {
		db.CreateQuestion(&question)
	}
	json.NewEncoder(w).Encode(message)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	params := mux.Vars(r)
	question := db.GetQuestion(params["id"])
	_ = json.NewDecoder(r.Body).Decode(&question)
	message, status := question.Validate()
	if status {
		db.UpdateQuestion(question)
	}
	json.NewEncoder(w).Encode(message)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	params := mux.Vars(r)
	question := db.GetQuestion(params["id"])
	db.DeleteQuestion(question)
	json.NewEncoder(w).Encode("success")
}
