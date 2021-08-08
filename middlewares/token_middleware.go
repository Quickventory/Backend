package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var hmacSampleSecret []byte

func ValidateTokenMiddleware(c *gin.Context) {
	_ = godotenv.Load(".env")
	fmt.Println([]byte(os.Getenv("JWT_SECRET")))
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

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
