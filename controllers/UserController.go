package controllers

import (
	"errors"
	"main/database"
	"main/models"
	"main/requests"
	"main/store"
	"main/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags users
// @Accept  json
// @Produce  json
// @Router /users/login [post]
func Login(c *gin.Context) {
	// get email and password from request
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

	tokenString := utils.GenerateTokenFromUser(&user, c)
	// log our user in the state and return success
	store.Store.User = user
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

	if userValidated, ok := utils.Validate(requests.RegisterAccountAndCompanyRequest{}, c); ok {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userValidated["password"]), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(400, gin.H{"message": "Exception occured"})
			return
		}
		// create new user
		user := models.User{
			Email:     userValidated["email"],
			Password:  string(hashedPassword),
			FirstName: userValidated["first_name"],
			LastName:  userValidated["last_name"],
		}

		database.Database.Create(&user)
		// create access token for user and save it in database and return it
		tokenString := utils.GenerateTokenFromUser(&user, c)

		store.Store.User = user

		c.JSON(200, gin.H{"token": tokenString})
	}

}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{"user": store.Store.User})
}
