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
    router.HandleFunc("/question/{id}", controllers.GetQuestion).Methods("GET")
    router.HandleFunc("/question", controllers.CreateQuestion).Methods("POST")
    router.HandleFunc("/question/{id}", controllers.UpdateQuestion).Methods("PUT")
    router.HandleFunc("/question/{id}", controllers.DeleteQuestion).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}
