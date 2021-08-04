package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidationStruct struct {
	toValidateStruct // here should store any struct given
}

func validate(class interface{}, c *gin.Context) (bool, error) {
	var result = ValidationStruct{toValidateStruct: class} //here any struct should be stored}
	var json = result.toValidateStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false, err
	}

	return true, nil
}
