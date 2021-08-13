package controllers

import (
	"errors"
	"fmt"
	"main/database"
	"main/models"
	"main/requests"
	"main/store"
	"main/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

func Register(c *gin.Context) {

	if userValidated, ok := utils.Validate(requests.UserCreateValidationRequest{}, c); ok {
		user := reflect.ValueOf(requests.UserCreateValidationRequest{}).Elem()
		n := user.FieldByName("Email").Interface().(string)
		fmt.Printf("%+v\n", n)

		fmt.Println(userValidated.(*requests.UserCreateValidationRequest).Email)
	}

}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{"user": store.Store.User})
}
