package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetQuizzes(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)

	if quizzes, err := models.GetQuizzes(); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	} else {
		json.NewEncoder(w).Encode(quizzes)
	}
}

func GetQuiz(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	param := mux.Vars(r)

	if quiz, err := models.GetQuiz(param["id"]); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	} else {
		json.NewEncoder(w).Encode(quiz)
	}
}

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	quiz := models.Quiz{}
	json.NewDecoder(r.Body).Decode(&quiz)

	if err := models.CreateQuiz(&quiz); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	} else {
		json.NewEncoder(w).Encode("Quiz created")
	}
}

func UpdateQuiz(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	params := mux.Vars(r)

	if quiz, err := models.GetQuiz(params["id"]); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	} else {
		json.NewDecoder(r.Body).Decode(&quiz)

		if err = models.UpdateQuiz(quiz); err != nil {
			json.NewEncoder(w).Encode(err)
			return
		} else {
			json.NewEncoder(w).Encode("Quiz updated")
		}
	}

}

func DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w)
	params := mux.Vars(r)

	if quiz, err := models.GetQuiz(params["id"]); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	} else {
		models.DeleteQuiz(quiz)
		json.NewEncoder(w).Encode("Quiz deleted")
	}

}
