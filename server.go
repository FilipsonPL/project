package main

import(
	"github.com/gin-gonic/gin"
	"github.com/FilipsonPL/project/controllers"
	"github.com/FilipsonPL/project/db"
)

func main() {
	ser := gin.Default()

	db.ConnectToDatabase()

	ser.POST("/register",Register)
	ser.GET("/show/users",ShowUsers)
	ser.POST("/unregister",Unregister)
	ser.POST("/edit/user",EditUser)

	ser.Run("localhost:8080")
}
