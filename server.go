package main

import(
	"github.com/gin-gonic/gin"
	"github.com/FilipsonPL/project/controllers"
	"github.com/FilipsonPL/project/db"
)

func main() {
	ser := gin.Default()

	db.ConnectToDatabase()

	ser.POST("/register",controllers.Register)
	ser.GET("/show/users",controllers.ShowUsers)
	ser.POST("/unregister",controllers.Unregister)
	ser.POST("/edit/user",controllers.EditUser)

	ser.Run("localhost:8080")
}
