package middlewares

import (
	"github.com/gin-gonic/gin"
)

func JsonMiddleware(c *gin.Context) {
	// set request header to "application/json"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Header("Content-Type", "application/json")
}
