package main

import (
	"JobBuddy/handlers"
	"JobBuddy/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {

	errEnvLoading := godotenv.Load()
	if errEnvLoading != nil {
		panic("Error loading .env file")
	}

	// Initialize Gin router

	router := gin.Default()

	// Define routes

	router.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": "pong",
		})

	})

	userApiGroup := router.Group("/api/user")
	{
		userApiGroup.POST("/register", handlers.HandleRegister)
		userApiGroup.GET("/email-confirm", handlers.HandleEmailConfirmation)
		userApiGroup.POST("/login", handlers.HandleLogin)
		userApiGroup.GET("/google-auth", handlers.HandleGoogleAuth)
		userApiGroup.GET("/google-auth-callback", handlers.HandleGoogleAuthCallback)

		userApiGroup.GET("/claims-checker", middlewares.Authenticator(), handlers.HandleClaimsChecker)

	}

	// Start the server

	router.Run(":8080")

}
