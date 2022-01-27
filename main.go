package main

import (
	"github.com/gin-gonic/gin"
	"go_jwt/controllers"
	middlewares "go_jwt/middleware"
	"go_jwt/models"
)

func main() {
	Router := gin.Default()
	models.InitialMigration()

	Router.POST("/register", controllers.Register)
	Router.POST("/login", controllers.Login)

	protected := Router.Group("/private")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/me", controllers.CurrentUser)

	Router.Run(":8000")
}
