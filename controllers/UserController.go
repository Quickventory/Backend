package controllers

import (
	"main/database"
	"main/models"
	"main/store"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterAccountAndCompanyRequest struct {
	Email     string `json:"email"  binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags users
// @Accept  json
// @Produce  json
// @Router /users/login [post]
func Login(c *gin.Context) {
}

// Register godoc
// @Summary Register a user
// @Description Register a user
// @Tags users
// @Accept  json
// @Produce  json
// @Router /users/register [post]
func Register(c *gin.Context) {
	var registrationData RegisterAccountAndCompanyRequest

	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registrationData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"message": "Exception occured"})
		return
	}
	// create new user
	user := models.User{
		Email:     registrationData.Email,
		Password:  string(hashedPassword),
		FirstName: registrationData.FirstName,
		LastName:  registrationData.LastName,
	}

	database.Database.Create(&user)
	// create access token for user and save it in database and return it
	tokenString := utils.GenerateTokenFromUser(&user, c)

	c.Set("user", user)

	c.JSON(200, gin.H{"token": tokenString})
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{"user": store.Store.User})
}
