package services

import (
	"JobBuddy/config"
	"JobBuddy/models/domain"
	"JobBuddy/models/dto"
	"JobBuddy/types"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func GetUser(field types.Field, value string) (*domain.User, error) {

	var user domain.User
	db, errDbAccess := config.AcessDB()

	if errDbAccess != nil {

		return nil, errDbAccess

	}

	var result *gorm.DB

	switch field {
	case types.ByID:
		result = db.Where("id = ?", value).First(&user)
		break
	case types.ByEmail:
		result = db.Where("email = ?", value).First(&user)
		break

	case types.ByEmailToken:
		result = db.Where("email_confirmation_token = ?", value).First(&user)
		break

	}

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

func UpdateUser(user *domain.User) error {

	db, errDbAccess := config.AcessDB()

	if errDbAccess != nil {

		return errDbAccess

	}

	result := db.Save(user)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func GenerateJWTToken(dto dto.UserLogin) (string, error) {

	var exp time.Time

	if dto.RememberMe {
		exp = time.Now().Add(time.Hour * 24)
	} else {
		exp = time.Now().Add(time.Hour)
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": dto.Email,
		"exp":   exp.Unix(),
	})

	secret := []byte(os.Getenv("JWT_SECRET"))

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
