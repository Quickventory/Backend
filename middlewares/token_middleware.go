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
	}

	requestToken = requestToken[7:] // This slice operation is to remove the "Bearer" string AND the space from the token

	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check if expiry is not passed
		if expiredAt, ok := claims["exp"]; ok {
			// check if exp is smaller than current time in unix
			if expiredAt.(int64) <= int64(time.Now().Unix()) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}

		// check if user_id claim is in the token
		if userId, ok := claims["user_id"]; ok {
			var accessToken models.AccessToken
			accessTokenRecord := database.Database.Model(&models.AccessToken{}).Where("user_id = ?", userId).First(&accessToken)
			err := accessTokenRecord.Error
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
			}

			toCompareToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user_id": accessToken.UserID,
			})

			signedToken, _ := toCompareToken.SignedString(hmacSampleSecret)

			if signedToken == requestToken {
				store.Store.User = accessToken.User
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}