package services

import (
	"JobBuddy/config"
	"JobBuddy/models/domain"
)

func GetUserByEmail(email string) (*domain.User, error) {

	var user domain.User
	db, errDbAccess := config.AcessDB()

	if errDbAccess != nil {

		return nil, errDbAccess

	}

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {

		println(result.Error.Error())

		return nil, result.Error
	}

	return &user, nil

}

func CreateUser(user *domain.User) error {

	db, errDbAccess := config.AcessDB()

	if errDbAccess != nil {

		return errDbAccess

	}

	result := db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil

}
