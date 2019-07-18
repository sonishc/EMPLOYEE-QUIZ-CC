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

func (question *Question) Validate() (string, bool) {
	if question.Text == "" {
		return "Question must have text", false
	}
	if question.Answer == "" {
		return "Question must have answer", false
	}
	return "Success", true
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

func GetQuestion(id string) *Question {
	question := &Question{}
	err := GetDB().Where("ID = ?", id).First(question).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return question
}

func CreateQuestion(question *Question) {
	GetDB().Create(question)
}

func UpdateQuestion(question *Question) {
	GetDB().Save(question)
}

func DeleteQuestion(question *Question) {
	GetDB().Delete(question)
}
