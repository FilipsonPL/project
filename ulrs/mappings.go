
package urls

import(
	"github.com/gin-gonic/gin"
	"github.com/FilipsonPL/project/controllers"
)

func SetUrls(server *Engine) {
	server.POST("/user", controllers.CreateUser)
	server.GET("/user/:id", controllers.FindUser)
	server.GET("/users", controllers.FindUsers)
	server.DELETE("/user/:id", controllers.DeleteUser)
	server.PUT("/user/:id", controllers.UpdateUser)
}