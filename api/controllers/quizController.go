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
    json.NewEncoder(w).Encode(models.GetQuizzes())
}

func GetQuiz(w http.ResponseWriter, r *http.Request) {
    utils.SetHeaders(w)
    param := mux.Vars(r)
    models.GetQuiz(param["id"])
    // quiz := models.GetQuiz(param["id"])

    // if status {
    //     json.NewEncoder(w).Encode(quiz)
    //     return
    // }
    // json.NewEncoder(w).Encode(errMessag)
}

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
    utils.SetHeaders(w)
    quiz := models.Quiz{}
    json.NewDecoder(r.Body).Decode(&quiz)

    err := models.CreateQuiz(&quiz)
    if err != nil {
        json.NewEncoder(w).Encode(err)
        return
    }
    json.NewEncoder(w).Encode("Success")
}

func UpdateQuiz(w http.ResponseWriter, r *http.Request) {
    utils.SetHeaders(w)
    params := mux.Vars(r)
    quiz := models.GetQuiz(params["id"])
    json.NewDecoder(r.Body).Decode(&quiz)

    err := models.UpdateQuiz(quiz)
    if err != nil {
        json.NewEncoder(w).Encode(err)
        return
    }
    json.NewEncoder(w).Encode("Success")
}

func DeleteQuiz(w http.ResponseWriter, r *http.Request) {
    utils.SetHeaders(w)
    params := mux.Vars(r)
    models.GetQuiz(params["id"])
    // quiz := models.GetQuiz(params["id"])

    // errMessag, status := quiz.Validate()
    // if status {
    //     models.DeleteQuiz(quiz)
    //     json.NewEncoder(w).Encode("Quiz deleted")
    //     return
    // }
    // json.NewEncoder(w).Encode(errMessag)
}
