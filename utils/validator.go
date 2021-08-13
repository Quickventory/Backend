package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(class interface{}, c *gin.Context) (interface{}, bool) {

	// check if the gin.Context body is not json or empty
	if c.ContentType() != "application/json" || c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return nil, false
	}

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, false
	}

	return &class, true
}
