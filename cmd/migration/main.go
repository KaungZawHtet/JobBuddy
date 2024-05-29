package main

import (
	"JobBuddy/config"

	"JobBuddy/models/domain"
	"github.com/joho/godotenv"
)

func main() {

	println("Hello From Migration")

	errEnvLoading := godotenv.Load()
	if errEnvLoading != nil {
		panic("Error loading .env file")
	}
	db, errAccessDB := config.AcessDB()

	if errAccessDB != nil {

		panic(errAccessDB.Error())
	}

	// Enable uuid-ossp extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	db.AutoMigrate(&domain.User{})

}
