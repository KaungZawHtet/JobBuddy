package main

import (
	"JobBuddy/config"
	"JobBuddy/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {

	errEnvLoading := godotenv.Load()
	if errEnvLoading != nil {
		panic("Error loading .env file")
	}

	db, errAccessDB := config.AcessDB()

	if errAccessDB != nil {

		panic(errAccessDB.Error())
	}

	// Initialize Gin router

	router := gin.Default()

	// Define routes

	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": db,
		})

	})

	userApiGroup := router.Group("/api/user")
	{
		userApiGroup.POST("/register", handlers.HandleRegister)

	}

	// Start the server

	router.Run(":8080")

}
