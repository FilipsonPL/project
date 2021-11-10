package main

import(
	"github.com/gin-gonic/gin"
	"github.com/FilipsonPL/project/db"
	"github.com/FilipsonPL/project/urls"
)

func main() {
	ser := gin.Default()

	db.ConnectToDatabase()

	ser.POST("/user", controllers.CreateUser)
	ser.GET("/user/:id", controllers.FindUser)
	ser.GET("/users", controllers.FindUsers)
	ser.DELETE("/user/:id", controllers.DeleteUser)
	ser.PUT("/user/:id", controllers.UpdateUser)

	ser.Run(":8080")
}
