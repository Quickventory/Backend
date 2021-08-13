package app

import (
	"fmt"
	"main/database"
	"main/middlewares"
	"main/routes"
	"main/store"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	store.Store = store.State{}

	fmt.Println("Initializing the application...")
	database.InitDatabase()
	router = gin.Default()
	router.Use(middlewares.JsonMiddleware)
	router.Use(middlewares.ValidateRequestMiddleware)
	routes.RegisterPublicRoutesForV1(router)
	routes.RegisterPrivateRoutesForV1(router)
}

// StartApp Start...
func StartApp() {
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
