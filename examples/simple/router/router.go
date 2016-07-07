package router

import (
	"github.com/wantedly/api-server-generator/examples/simple/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("api")
	{
		//USER API
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}
}
