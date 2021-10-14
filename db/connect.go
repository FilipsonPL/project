package db

import (
	"github.com/FilipsonPL/project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase() {
	Database, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	Database.AutoMigrate(&models.User{})
}