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

	log.Fatal(http.ListenAndServe(":8000", router))
}
