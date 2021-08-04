package app

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func mapUrls() {
	mapUserUrls()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})
}

func mapUserUrls() *gin.RouterGroup {

	usersGroup := router.Group("api/users")
	{
		usersGroup.POST("/login", controllers.Login)
		usersGroup.POST("/register", controllers.Register)
	}

	return usersGroup
}
