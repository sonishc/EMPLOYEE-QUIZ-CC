package db

import (
	"../models"
	"fmt"
)

func SeedQuestions() {
	questions := make([]*models.Question, 0)

	GetDB().Create(&models.Question{Text: "5+6", Answer: "11"})
	GetDB().Create(&models.Question{Text: "2+7", Answer: "9"})
	GetDB().Create(&models.Question{Text: "1+6", Answer: "7"})
	GetDB().Create(&models.Question{Text: "7+9", Answer: "16"})
	GetDB().Create(&models.Question{Text: "3+5", Answer: "8"})


	GetDB().Find(&questions)
	for _,value := range questions {
		fmt.Println(*value)
	}
}

func GetQuestions() []*models.Question {
	questions := make([]*models.Question, 0)
	err := GetDB().Find(&questions).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return questions
}
