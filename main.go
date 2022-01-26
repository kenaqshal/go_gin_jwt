package main

import (
	"github.com/gin-gonic/gin"
	"go_jwt/controllers"
	"go_jwt/models"
)

func main() {
	Router := gin.Default()
	models.InitialMigration()

	Router.POST("/sign-up", controllers.SignUp)
	Router.POST("/sign-in", controllers.SignIn)

	authorized := Router.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	Router.Run(":8000")
}
