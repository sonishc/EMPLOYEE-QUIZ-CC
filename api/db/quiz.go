// package models

// import (
// 	"../models"
// 	"fmt"
// )

// func GetQuizzes() []*models.Quiz {
// 	quizzes := make([]*models.Quiz, 0)
// 	err := GetDB().Find(&quizzes).Error

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	return quizzes
// }

// func GetQuiz(id string) *models.Quiz {
// 	quiz := &models.Quiz{}
// 	err := GetDB().First(&quiz, id).Error

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return quiz
// }

// func CreateQuiz(quiz *models.Quiz) {
// 	GetDB().Create(quiz)
// }

// func UpdateQuiz(quiz *models.Quiz) {
// 	GetDB().Save(quiz)
// }

// func DeleteQuiz(quiz *models.Quiz) {
// 	GetDB().Delete(quiz)
// }
