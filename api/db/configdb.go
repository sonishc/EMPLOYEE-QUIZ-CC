// package db

// import (
// 	"../models"
// 	"fmt"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// )

// var db *gorm.DB

// func init() {
// 	connection, e := gorm.Open("postgres", "user=quiz password=quiz dbname=quiz")
// 	if e != nil {
// 		fmt.Print(e)
// 		panic("Could not connect to database")
// 	}

// 	db = connection
// 	db.AutoMigrate(
// 		&models.Question{},
// 		&models.Quiz{},
// 	)
// }

// func GetDB() *gorm.DB {
// 	return db
// }
