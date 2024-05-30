package main

import (
	"JobBuddy/handlers"
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

	// Define routes

	router.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": "pong",
		})

	})

	userApiGroup := router.Group("/api/user")
	{
		userApiGroup.POST("/register", handlers.HandleRegister)

	}

	// Start the server

	router.Run(":80")

}
