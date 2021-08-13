package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(class interface{}, c *gin.Context) (map[string]string, bool) {

	// check if the gin.Context body is not json or empty
	if c.ContentType() != "application/json" || c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return nil, false
	}

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, false
	}

	// create a map[string]string to hold the class values
	m := make(map[string]string)
	// iterate over the map[string]interface{}
	for key, value := range class.(map[string]interface{}) {
		// append key value of class to m
		m[key] = value.(string)
	}

	return m, true
}
