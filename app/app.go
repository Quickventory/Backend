package app

import (
	"fmt"
	"main/database"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	fmt.Println("Initializing the application...")
	database.InitDatabase()
	router = gin.Default()

}

// StartApp Start...
func StartApp() {
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
