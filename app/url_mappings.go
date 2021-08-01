package app

import (
	"main/controllers"
)

func mapUrls() {
	mapUserUrls()
}

func mapUserUrls() {

	usersGroup := router.Group("api/users")
	{
		usersGroup.POST("/login", controllers.Login)
		usersGroup.POST("/register", controllers.Register)
	}
	// router.GET("api/users/:user_id", controllers.GetUser)
}
