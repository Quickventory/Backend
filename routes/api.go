package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutesForV1(router *gin.Engine) {
	// create an api group
	router.GET("/")
	apiGroup := router.Group("/api")
	v1group := apiGroup.Group("/v1")
	usersGroup := v1group.Group("/users")
	{
		usersGroup.POST("/login", controllers.Login)
		usersGroup.POST("/register", controllers.Register)
	}
}
