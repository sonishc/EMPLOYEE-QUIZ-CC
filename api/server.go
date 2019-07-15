package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

type Question struct {
  ID     string `json:"id"`
  Text   string `json:"text"`
  Answer string `json:"answer"`
}

var questions []Question

func seedQuestions() {
  questions = append(questions, Question{ID: "1", Text: "5+6", Answer: "11"})
  questions = append(questions, Question{ID: "2", Text: "2+7", Answer: "9"})
  questions = append(questions, Question{ID: "3", Text: "1+6", Answer: "7"})
  questions = append(questions, Question{ID: "4", Text: "7+9", Answer: "16"})
  questions = append(questions, Question{ID: "5", Text: "3+5", Answer: "8"})
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
  json.NewEncoder(w).Encode(questions)
}

func main() {
  router := mux.NewRouter()
  seedQuestions()
  router.HandleFunc("/", getQuestions).Methods("GET")

  log.Fatal(http.ListenAndServe(":8000", router))
}
