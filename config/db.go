package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func AcessDB() (*gorm.DB, error) {

	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")

	dbUserName := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbPort := os.Getenv("DATABASE_PORT")

	log.Printf("Connecting to database with: host=%s user=%s dbname=%s port=%s password=%s", dbHost, dbUserName, dbName, dbPort, dbPassword)

	//check ssl mode at https://gorm.io/docs/connecting_to_the_database.html
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Bangkok", dbHost, dbUserName, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db, nil

}
