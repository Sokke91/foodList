package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	db, err := gorm.Open(sqlite.Open("foodListDB"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect with the database")
	}
	DB = db
}
