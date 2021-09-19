package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/FilipsonPL/project/models"
)

var Database *gorm.DB

func ConnectToDatabase() {
	connection, err := gorm.Open("sqlite3", "students.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	connection.AutoMigrate(&models.User{})

	Database = connection
}