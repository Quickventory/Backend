package controllers

import (
	"errors"
	"main/database"
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	var login LoginRequest
	var user models.User
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	record := database.Database.Model(&models.User{}).Where("email = ?", login.Email).First(&user)
	recordErr := record.Error
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)

	if err != nil || errors.Is(recordErr, gorm.ErrRecordNotFound) {
		// return invalid credentials error
		c.JSON(400, gin.H{"message": "Invalid credentials"})
		return
	}

	// check if password matches
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(user.Password)); err != nil {
		// return invalid credentials error
		c.JSON(400, gin.H{"message": "Invalid credentials"})
		return
	}

	tokenString := utils.GenerateTokenFromUser(&user, c)
	// log our user in the state and return success
	c.Set("user", user)
	// return token
	c.JSON(200, gin.H{"token": tokenString})
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
		return
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

	dbErr := database.Database.Create(&user)

	// look for duplicate key value string in error
	if err := dbErr.Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// create access token for user and save it in database and return it
	tokenString := utils.GenerateTokenFromUser(&user, c)

	c.Set("user", user)

	c.JSON(200, gin.H{"token": tokenString})
}

func GetUser(c *gin.Context) {
	// c.JSON(200, gin.H{"user": c.Get("user")Ç
}
