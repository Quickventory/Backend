package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"main/database"
	"main/models"
	"main/store"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var hmacSampleSecret []byte

func ValidateTokenMiddleware(c *gin.Context) {
	_ = godotenv.Load(".env")
	hmacSampleSecret = []byte(os.Getenv("JWT_SECRET"))
	// get beaer token from authorization header
	requestToken := c.Request.Header.Get("Authorization")

	if requestToken == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	requestToken = requestToken[7:] // This slice operation is to remove the "Bearer" string AND the space from the token

	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check if expiry is not passed
		if expiredAt, ok := claims["exp"]; ok {
			// check if exp is smaller than current time in unix
			if expiredAt.(float64) <= float64(time.Now().Unix()) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		// check if user_id claim is in the token
		if userId, ok := claims["user_id"]; ok {
			var accessToken models.AccessToken
			accessTokenRecord := database.Database.Model(&models.AccessToken{}).Where("user_id = ?", userId).First(&accessToken)
			err := accessTokenRecord.Error
			if err != nil {
				c.AbortWithStatus(404)
				return
			}

			toCompareToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user_id": accessToken.UserID,
				"exp":     accessToken.ExpiresAt.Unix(),
				"iss":     accessToken.IssuedAt.Unix(),
			})

			signedToken, _ := toCompareToken.SignedString(hmacSampleSecret)
			var user models.User
			userRecord := database.Database.Model(&models.User{}).Where("id = ?", accessToken.UserID).First(&user)
			err = userRecord.Error
			if err != nil {
				c.AbortWithStatus(404)
				return
			}
			if signedToken == requestToken {
				store.Store.User = user
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
