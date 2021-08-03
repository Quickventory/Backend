package app

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	mapUserUrls()
}

func mapUserUrls() *gin.RouterGroup {

	usersGroup := router.Group("api/users")
	{
		usersGroup.POST("/login", controllers.Login)
		usersGroup.POST("/register", controllers.Register)
	}

	return usersGroup
}
