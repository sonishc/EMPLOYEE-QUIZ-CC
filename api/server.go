package main

import (
	"../api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Question router
	router.HandleFunc("/", controllers.GetQuestions).Methods("GET")
	router.HandleFunc("/question/{id}", controllers.GetQuestion).Methods("GET")
	router.HandleFunc("/question", controllers.CreateQuestion).Methods("POST")
	router.HandleFunc("/question/{id}", controllers.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/question/{id}", controllers.DeleteQuestion).Methods("DELETE")

	// Quiz router
	router.HandleFunc("/quizzes", controllers.GetQuizzes).Methods("GET")
	router.HandleFunc("/quizzes/{id}", controllers.GetQuiz).Methods("GET")
	router.HandleFunc("/quizzes", controllers.CreateQuiz).Methods("POST")
	router.HandleFunc("/quizzes/{id}", controllers.UpdateQuiz).Methods("PUT")
	router.HandleFunc("/quizzes/{id}", controllers.DeleteQuiz).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
