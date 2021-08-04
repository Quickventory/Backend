package app

import (
	"main/controllers"
	"main/requests"
	"main/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// create an associative array of names that hold a map of the gin.context method and its request validation
var requestValidationName = map[string]map[string]interface{}{
	"users": {
		"POST": requests.UserCreateValidationRequest{},
		"GET":  requests.UserCreateValidationRequest{},
	},
}

func validateRequestHandler(c *gin.Context) {
	// check if gin.context method is "POST", "PUT", "PATCH"
	if condition := c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH"; !condition {
		//loop over gin.Engine registered routes
		for _, route := range router.Routes() {
			// split the string of the route.Path to get the second part(the model)
			path := strings.Split(route.Path, "/")[2]
			// check if the route.Path is in the requestValidationName map
			if _, ok := requestValidationName[path]; ok {
				//check if the gin.context method is in the requestValidationName map
				if class, ok := requestValidationName[path][c.Request.Method]; ok {
					if ok := utils.Validate(class, c); ok {
						c.Next()
					}
				}
			}
		}
	}
}

func mapUrls() {
	mapUserUrls()
	router.GET("/api/users/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})
}

func mapUserUrls() *gin.RouterGroup {

	usersGroup := router.Group("api/users")
	{
		usersGroup.POST("/login", controllers.Login)
		// usersGroup.GET("/", controllers.Login)
		usersGroup.POST("/register", controllers.Register)
	}

	return usersGroup
}
