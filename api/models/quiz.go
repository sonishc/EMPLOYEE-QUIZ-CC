package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Quiz struct {
	gorm.Model
	Title string `json:title"`
}

// VALIDATIONS

type QuizValidationError struct {
	WhatHappened []string `json:WhatHappened"`
}

func (e QuizValidationError) Error() string {
	return fmt.Sprintf("Quiz: %v", e.WhatHappened)
}

func (quiz *Quiz) Validate() ([]string, bool) {
	errMessages := make([]string, 0)

	if quiz.isPresenceTitle() == false {
		errMessages = append(errMessages, "Quiz must have title")
	}
	if quiz.isUniqTitle() == false {
		errMessages = append(errMessages, "Title must have be uniq")
	}
	status := len(errMessages) == 0

	return errMessages, status
}

// Checks if the title is present
func (quiz *Quiz) isPresenceTitle() bool {
	if quiz.Title == "" {
		return false
	}
	return true
}

// Checks whether the title is unique
func (quiz *Quiz) isUniqTitle() bool {
	quizTemplate := &Quiz{}

	if result := db.Where("title = ?", quiz.Title).First(&quizTemplate); result.Error != nil {
		return true
	}
	return false
}

// END VALIDATIONS

// Select all quizzes
func GetQuizzes() ([]*Quiz, error) {
	quizzes := make([]*Quiz, 0)

	if result := GetDB().Find(&quizzes); result.Error != nil {
		return nil, QuizValidationError{[]string{"Not found"}}
	}
	return quizzes, nil
}

// Select one quiz
func GetQuiz(id string) (*Quiz, error) {
	quiz := &Quiz{}

	if result := db.Where("id = ?", id).First(&quiz); result.Error != nil {
		return nil, QuizValidationError{[]string{"Not found"}}
	}
	return quiz, nil
}

// Create one quiz
func CreateQuiz(quiz *Quiz) error {
	msg, status := quiz.Validate()
	err := QuizValidationError{msg}

	if status {
		GetDB().Create(quiz)
		return nil
	}
	return err
}

// Update one quiz
func UpdateQuiz(quiz *Quiz) error {
	msg, status := quiz.Validate()
	err := QuizValidationError{msg}

	if status {
		GetDB().Save(quiz)
		return nil
	}
	return err
}

// Delete one quiz
func DeleteQuiz(quiz *Quiz) {
	GetDB().Delete(quiz)
}
