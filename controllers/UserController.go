package controllers

import (
	"errors"
	"main/database"
	"main/models"
	"main/store"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	// get email and password from request
	_ = godotenv.Load(".env")
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))
	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")

	record := database.Database.Model(&models.User{}).Where("email = ?", email).First(&user)
	recordErr := record.Error
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

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

	// create a token using claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     int64(time.Now().Add(time.Hour * 24).Unix()),
		"iss":     int64(time.Now().Unix()),
	})

	// create a new access token record and save it
	accessToken := models.AccessToken{
		UserID:    user.ID,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	database.Database.Create(&accessToken)
	// sign token
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid credentials"})
		return
	}

	// log our user in the state and return success
	store.Store.User = user
	// return token
	c.JSON(200, gin.H{"token": tokenString})

}

func Register(c *gin.Context) {

}
