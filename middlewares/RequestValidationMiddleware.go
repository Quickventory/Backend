package middlewares

import (
	"main/requests"
	"main/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var requestValidationName = map[string]map[string]interface{}{
	"users/register": {
		"POST": requests.RegisterAccountAndCompanyRequest{},
	},
}

func ValidateRequestMiddleware(c *gin.Context) {
	// check if gin.context method is "POST", "PUT", "PATCH"
	if condition := c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH"; condition {
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
		} else if len(pathParts) == 3 {
			//get our third element of pathParts which will be our model name
			path = pathParts[2]

			// remove our fucking "/" from the path
			path = strings.TrimRight(path, "/")
		} else {
			// if we don't have a valid path, we'll just return an error
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		// check if the route path is in our requestValidationName map
		if _, ok := requestValidationName[path]; ok {
			// check if the current request method is in our requestValidationName map
			if class, ok := requestValidationName[path][c.Request.Method]; ok {
				if _, ok := utils.Validate(class, c); ok {
					c.Next()
				}
			}
		}
	}

	c.Next()
}
