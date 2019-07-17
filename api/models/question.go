package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Text        string `json:"text"`
	Answer      string `json:"answer"`
}
