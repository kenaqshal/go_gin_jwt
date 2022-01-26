package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuthorized(handler http.HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {

		if c.Request.Header["Token"] == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password Incorrect!"})
			return
		}

		var mySigningKey = []byte("jwt_go")

		token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Your token has been expired!"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				c.Request.Header.Set("Role", "admin")
				return

			} else if claims["role"] == "user" {

				c.Request.Header.Set("Role", "user")
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not Authorized!"})
		return
	}
}
