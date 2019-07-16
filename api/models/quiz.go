package models

type Quiz struct {
	gorm.Model
	Title string `json:title"`
}
