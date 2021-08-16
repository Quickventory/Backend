package app

import (
	"fmt"
	"main/database"
	"main/routes"
	"main/store"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	router *gin.Engine
)

func init() {
	store.Store = store.State{}

	fmt.Println("Initializing the application...")
	database.InitDatabase()
	router = gin.Default()
	routes.RegisterPublicRoutesForV1(router)
	routes.RegisterPrivateRoutesForV1(router)
}

// StartApp Start...
func StartApp() {
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
