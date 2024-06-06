package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func AcessDB() (*gorm.DB, error) {

	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")

	dbUserName := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbPort := os.Getenv("DATABASE_PORT")

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	//check ssl mode at https://gorm.io/docs/connecting_to_the_database.html
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Bangkok", dbHost, dbUserName, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic(err.Error())
	}

	return db, nil

}
