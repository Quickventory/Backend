package app

import (
	"fmt"
	"main/database"
	"main/middlewares"
	"main/routes"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	fmt.Println("Initializing the application...")
	database.InitDatabase()
	router = gin.Default()
	router.Use(middlewares.ValidateRequestMiddleware)
	routes.RegisterRoutes(router)
}

// StartApp Start...
func StartApp() {
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
