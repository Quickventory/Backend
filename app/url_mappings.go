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
	"users/foo": {
		"POST": requests.UserCreateValidationRequest{},
		"GET":  requests.UserCreateValidationRequest{},
	},

	"foo": {
		"GET": requests.UserCreateValidationRequest{},
	},
}

func validateRequestHandler(c *gin.Context) {
	// check if gin.context method is "POST", "PUT", "PATCH"
	// DEV: "GET" is a valid method, but we don't want to validate it except for now, hence why we have !condition and not condition
	if condition := c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH"; !condition {
		//get the current route path
		routePath := c.Request.URL.Path
		// split the routepath into a slice of strings
		pathParts := strings.Split(routePath, "/")
		// check if len(pathParts) > 3
		path := ""
		if len(pathParts) > 3 {
			// loop over the path parts and concatenate them starting at the second element of our pathParts slice
			for i := 2; i < len(pathParts); i++ {
				path += pathParts[i] + "/"
			}
			// remove the last "/" from the path
			path = path[:len(path)-1]
		} else {
			//get our third element of pathParts which will be our model name
			path = pathParts[2]

			// remove our fucking "/" from the path
			path = strings.TrimRight(path, "/")

		}

		// check if the route path is in our requestValidationName map
		if _, ok := requestValidationName[path]; ok {
			// check if the current request method is in our requestValidationName map
			if class, ok := requestValidationName[path][c.Request.Method]; ok {
				if ok := utils.Validate(class, c); ok {
					c.Next()
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
		usersGroup.GET("/foo", controllers.Login)
		// usersGroup.GET("/", controllers.Login)
		usersGroup.POST("/register", controllers.Register)
	}

	return usersGroup
}
