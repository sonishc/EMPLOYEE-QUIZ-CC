package main

import (
	"../api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetQuestions).Methods("GET")

	// Quiz router
	router.HandleFunc("/quizzes", controllers.GetQuizzes).Methods("GET")
	router.HandleFunc("/quizzes/{id}", controllers.GetQuiz).Methods("GET")
	router.HandleFunc("/quizzes", controllers.CreateQuiz).Methods("POST")
	router.HandleFunc("/quizzes/{id}", controllers.UpdateQuiz).Methods("PUT")
	router.HandleFunc("/quizzes/{id}", controllers.DeleteQuiz).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
