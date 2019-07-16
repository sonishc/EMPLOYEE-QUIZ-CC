package main

import (
  "../api/controllers"
  "../api/models"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

func main() {
  router := mux.NewRouter()
  models.SeedQuestions()
  router.HandleFunc("/", controllers.GetQuestions).Methods("GET")

  log.Fatal(http.ListenAndServe(":8000", router))
}
