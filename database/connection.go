package database

import(
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase(filename string) {
	connection, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})

	if err!=nil {
		panic("Failure during connecting to database!")
	}

	Database = connection
	connection.AutoMigrate(&models.User{})
}