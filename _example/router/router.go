package router

import (
	"github.com/wantedly/apig/_example/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("api")
	{

		api.GET("/companies", controllers.GetCompanies)
		api.GET("/companies/:id", controllers.GetCompany)
		api.POST("/companies", controllers.CreateCompany)
		api.PUT("/companies/:id", controllers.UpdateCompany)
		api.DELETE("/companies/:id", controllers.DeleteCompany)

		api.GET("/emails", controllers.GetEmails)
		api.GET("/emails/:id", controllers.GetEmail)
		api.POST("/emails", controllers.CreateEmail)
		api.PUT("/emails/:id", controllers.UpdateEmail)
		api.DELETE("/emails/:id", controllers.DeleteEmail)

		api.GET("/jobs", controllers.GetJobs)
		api.GET("/jobs/:id", controllers.GetJob)
		api.POST("/jobs", controllers.CreateJob)
		api.PUT("/jobs/:id", controllers.UpdateJob)
		api.DELETE("/jobs/:id", controllers.DeleteJob)

		api.GET("/profiles", controllers.GetProfiles)
		api.GET("/profiles/:id", controllers.GetProfile)
		api.POST("/profiles", controllers.CreateProfile)
		api.PUT("/profiles/:id", controllers.UpdateProfile)
		api.DELETE("/profiles/:id", controllers.DeleteProfile)

		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

	}
}
