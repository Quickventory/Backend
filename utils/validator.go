package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(class interface{}, c *gin.Context) bool {
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	return true
}
