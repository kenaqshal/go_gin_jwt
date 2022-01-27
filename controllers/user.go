package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_jwt/helper"
	"go_jwt/models"
	"net/http"
)

func CurrentUser(c *gin.Context) {

	email, err := helper.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := GetUserByID(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
func GetUserByID(email interface{}) (models.User, error) {
	connection := models.GetDatabase()
	defer models.Closedatabase(connection)
	var user models.User

	if err := connection.Where("email = ?", email).First(&user).Error; err != nil {
		return user, errors.New("User not found!")
	}

	return user, nil

}
