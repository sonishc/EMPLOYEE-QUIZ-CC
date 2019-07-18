package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Quiz struct {
	gorm.Model
	Title string `json:title"`
}

// Validations

// func (quiz *Quiz) Validate() (string, bool) {
// 	switch {
// 	case quiz.isPresenceTitle() == false:
// 		msg := "Quiz must have title"
// 		return msg, false
// 	case quiz.isUniqTitle() == false:
// 		msg := "Title not uniq"
// 		return msg, false
// 	default:
// 		return "OK", true
// 	}
// }

func (quiz *Quiz) isPresenceTitle() bool {
	if quiz.Title == "" {
		return false
	}
	return true
}

// func (quiz *Quiz) isUniqTitle() bool {
// 	// err := GetDB().First(quiz).Error
// 	// if err != nil {
// 	// 	return false
// 	// }
// 	// return true
// 	return true
// }

// DB queries

func GetQuizzes() []*Quiz {
	quizzes := make([]*Quiz, 0)
	err := GetDB().Find(&quizzes).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return quizzes
}

func GetQuiz(id string) *Quiz {
	quiz := &Quiz{}
	err := GetDB().First(quiz, id).Error

	if err != nil {
		fmt.Println(err)
	}
	return quiz
}

func CreateQuiz(quiz *Quiz) error {
	valid := quiz.isPresenceTitle()

	if valid {
		GetDB().Create(quiz)
		return nil
	}
	return errors.New("")
}

func UpdateQuiz(quiz *Quiz) {
	valid := quiz.isPresenceTitle()

	if valid {
		GetDB().Save(quiz)
		return nil
	}
	return errors.New("")
}

func DeleteQuiz(quiz *Quiz) {
	GetDB().Delete(quiz)
}
