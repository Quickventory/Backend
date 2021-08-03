package controllers

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {
	db := database.Database

	db.Find(c, "users", "email")
}
