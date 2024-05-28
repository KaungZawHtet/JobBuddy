package main

import (
	"JobBuddy/config"
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

	db, err := config.AcessDB()

	if err != nil {

		panic(err.Error())
	}

	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{

			"message": db,
		})

	})

	// Start the server

	router.Run(":8080")

}
