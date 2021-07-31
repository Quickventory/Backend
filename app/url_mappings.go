package app

import "main/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
