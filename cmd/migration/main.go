package main

import (
	"JobBuddy/config"
	"JobBuddy/models/domain"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func main() {

	println("Hello From Migration")

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	// Construct the absolute path to the .env file
	envPath := filepath.Join(wd, ".env")

	// Load the .env file
	errEnvLoading := godotenv.Load(envPath)
	if errEnvLoading != nil {
		log.Fatalf("Error loading .env file: %v", errEnvLoading)
	}
	db, errAccessDB := config.AcessDB()

	if errAccessDB != nil {

		panic(errAccessDB.Error())
	}

	// Enable uuid-ossp extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.JobApplication{})

}
