package services

import (
	"JobBuddy/config"
	"JobBuddy/models/domain"
	"JobBuddy/models/dto"
	"JobBuddy/types"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
		"id":    dto.Id,
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

func ValidateJWTToken(tokenString string) (jwt.MapClaims, bool, error) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := []byte(os.Getenv("JWT_SECRET"))

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		return nil, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	return claims, ok, nil

}

func ExchangeCodeForGoogleToken(code string) (*oauth2.Token, error) {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_AUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_AUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"),
		Endpoint:     google.Endpoint,
	}

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func FetchGoogleUserProfile(accessToken string) (dto.UserProfile, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo", nil)
	if err != nil {
		return dto.UserProfile{}, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return dto.UserProfile{}, err
	}
	defer resp.Body.Close()

	var profile dto.UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return dto.UserProfile{}, err
	}

	return profile, nil
}

func GetClaimVar(by types.Field, context *gin.Context) (interface{}, error) {

	claims, exists := context.Get("mapClaims")

	if !exists {

		return "", fmt.Errorf("unauthorization detected")
	}

	mapClaims, ok := claims.(jwt.MapClaims)

	if !ok {
		return "", fmt.Errorf("claims error")
	}

	switch by {
	case types.ByEmail:

		return mapClaims["email"].(string), nil

	case types.ByID:

		return mapClaims["id"].(uuid.UUID), nil

	}
	return "", fmt.Errorf("some err happened")
}
