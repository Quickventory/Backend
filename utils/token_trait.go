package utils

import (
	"main/database"
	"main/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateTokenFromUser(user *models.User, c *gin.Context) string {
	_ = godotenv.Load(".env")
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iss":     time.Now().Unix(),
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
		c.JSON(400, gin.H{"message": "Exception occured"})
	}

	return tokenString
}
