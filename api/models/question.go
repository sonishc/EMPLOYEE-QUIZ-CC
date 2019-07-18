package models

import (
    "fmt"
    "github.com/jinzhu/gorm"
)

type Question struct {
    gorm.Model
    Text   string `json:"text"`
    Answer string `json:"answer"`
}

func GetQuestions() []*Question {
    questions := make([]*Question, 0)
    err := GetDB().Find(&questions).Error
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return questions
}
