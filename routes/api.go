package routes

import (
	"main/controllers"
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutesForV1(router *gin.Engine) {
	// create an api group
	public := router.Group("/api/v1")
	usersGroup := public.Group("/users")
	{
		usersGroup.POST("/login", controllers.Login)
		usersGroup.POST("/register", controllers.Register)
	}
}

func RegisterPrivateRoutesForV1(router *gin.Engine) {
	// create an private api group
	private := router.Group("/api/v1")
	private.Use(middlewares.ValidateTokenMiddleware)
	usersGroup := private.Group("/users")
	{
		usersGroup.GET("/me", controllers.GetUser)
	}
}
