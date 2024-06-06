package main

import (
	"JobBuddy/handlers"
	"JobBuddy/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	errEnvLoading := godotenv.Load()
	if errEnvLoading != nil {
		panic("Error loading .env file")
	}

	// Initialize Gin router

	router := gin.Default()

	router.Use(gin.Recovery())

	// Define routes

	router.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": "pong",
		})

	})

	userApiGroup := router.Group("/api/users")
	{
		userApiGroup.POST("/register", handlers.HandleRegister)
		userApiGroup.GET("/email-confirm", handlers.HandleEmailConfirmation)
		userApiGroup.POST("/login", handlers.HandleLogin)
		userApiGroup.GET("/google-auth", handlers.HandleGoogleAuth)
		userApiGroup.GET("/google-auth-callback", handlers.HandleGoogleAuthCallback)

		userApiGroup.GET("/claims-checker", middlewares.Authenticator(), handlers.HandleClaimsChecker)

	}

	JobApplicationApiGroup := router.Group("/api/job-applications")
	{

		JobApplicationApiGroup.GET("", middlewares.Authenticator(), handlers.HandleMyApplicationsList)
		JobApplicationApiGroup.POST("", middlewares.Authenticator(), handlers.HandleMyApplicationCreation)
		JobApplicationApiGroup.DELETE("/:id", middlewares.Authenticator(), handlers.HandleMyJobApplicationDeletion)

	}

	// Start the server

	router.Run(":8080")

}
