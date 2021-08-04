package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(class interface{}, c *gin.Context) (bool, error) {
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false, err
	}

	return true, nil
}
