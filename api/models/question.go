package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Text        string `json:"text"`
	Answer      string `json:"answer"`
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
