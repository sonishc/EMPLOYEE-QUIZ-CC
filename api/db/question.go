package db

import (
	"../models"
	"fmt"
)

func GetQuestions() []*models.Question {
	questions := make([]*models.Question, 0)
	err := GetDB().Find(&questions).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return questions
}

func GetQuestion(id string) *models.Question {
	question := &models.Question{}
	err := GetDB().Where("ID = ?", id).First(question).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return question
}

func CreateQuestion(question *models.Question) {
	GetDB().Create(question)
}

func UpdateQuestion(question *models.Question) {
	GetDB().Save(question)
}

func DeleteQuestion(question *models.Question) {
	GetDB().Delete(question)
}
