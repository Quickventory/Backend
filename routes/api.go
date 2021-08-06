package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// create an api group
	apiGroup := router.Group("/api")
	usersGroup := apiGroup.Group("/users")
	{
		usersGroup.POST("/login", controllers.Login)
		usersGroup.POST("/register", controllers.Register)
	}
}
