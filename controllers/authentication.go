package controllers

import (
	"github.com/gin-gonic/gin"
	"go_jwt/helper"
	"go_jwt/models"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	connection := models.GetDatabase()
	defer models.Closedatabase(connection)

	var authdetails models.Authentication
	if err := c.ShouldBindJSON(&authdetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var authuser models.User
	if error := connection.Where("email = ?", authdetails.Email).First(&authuser).Error; error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email Not Found!"})
		return
	}

	check := helper.CheckPasswordHash(authdetails.Password, authuser.Password)

	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password Incorrect!"})
		return
	}

	validToken, err := helper.GenerateJWT(authuser.Email, authuser.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password Incorrect!"})
		return
	}

	var token models.Token
	token.Email = authuser.Email
	token.Role = authuser.Role
	token.TokenString = validToken

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully!",
		"data":    token,
	})
	return

}

func Register(c *gin.Context) {
	connection := models.GetDatabase()
	defer models.Closedatabase(connection)

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//var userData models.User

	//checks if email is already register or not
	if error := connection.Where("email = ?", input.Email).First(&input).Error; error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered!"})
		return
	}

	var err error
	input.Password, err = helper.GeneratehashPassword(input.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//insert user details in database
	connection.Create(&input)
	c.JSON(http.StatusOK, gin.H{"message": "Success sign up!"})

}
