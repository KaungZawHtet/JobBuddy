package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	//"JobBuddy/config"
	"JobBuddy/models/domain"
	"JobBuddy/models/dto"
	"JobBuddy/services"
	"JobBuddy/types"

	"golang.org/x/crypto/bcrypt"

	"net/http"
	"os"
)

func HandleRegister(context *gin.Context) {

	var userRegistration dto.UserRegistration

	errBindJson := context.ShouldBindJSON(&userRegistration)

	if errBindJson != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": errBindJson.Error(),
		})

		return
	}

	if userRegistration.Email != userRegistration.EmailRecheck {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Email should be same",
		})
		return
	}

	checkedUser, _ := services.GetUser(types.ByEmail, userRegistration.Email)

	if checkedUser != nil && checkedUser.ID.String() != "" {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Email has already existed",
			//"data":         *checkedUser,
			//"internal err": errGetUser.Error(),
		})

		return
	}

	//create email token

	emailToken, errToken := services.GenerateEmailToken()

	if errToken != nil {

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errToken.Error(),
			//"data":    user,
		})
		return

	}

	hashedPassword, errPwHash := bcrypt.GenerateFromPassword([]byte(userRegistration.Password), bcrypt.DefaultCost)

	if errPwHash != nil {

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errPwHash.Error(),
			//"data":    user,
		})
		return

	}

	newUser := domain.User{
		UserName:               userRegistration.UserName,
		FirstName:              userRegistration.FirstName,
		LastName:               userRegistration.LastName,
		Email:                  userRegistration.Email,
		EmailConfirmationToken: emailToken,
		Password:               string(hashedPassword),
	}

	//TODO: send email confirmation letter

	errCreateUser := services.CreateUser(&newUser)

	if errCreateUser != nil {

		context.JSON(http.StatusAccepted, gin.H{
			"message": errCreateUser.Error(),
		})

		return

	}

	_, errorEmail := services.SendEmailConfirmationLink(newUser.Email, emailToken)

	if errorEmail != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errorEmail.Error(),
		})

		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Check your email",
		"data":    newUser,
	})

}

func HandleEmailConfirmation(context *gin.Context) {

	token := context.Query("token")

	checkedUser, _ := services.GetUser(types.ByEmailToken, token)

	if checkedUser != nil && checkedUser.EmailConfirmationToken == "" {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error at confirmation",
		})

		return
	}

	checkedUser.EmailConfirmed = true

	services.UpdateUser(checkedUser)

	context.JSON(http.StatusOK, gin.H{
		"message": "Email Confirmed! Please Log in with your email",
	})

}

func HandleLogin(context *gin.Context) {

	var loginForm dto.UserLogin

	errBindJsong := context.ShouldBindJSON(&loginForm)

	if errBindJsong != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": errBindJsong.Error(),
		})
		return
	}

	checkedUser, _ := services.GetUser(types.ByEmail, loginForm.Email)

	if checkedUser != nil && checkedUser.Email == "" {

		context.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return

	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(checkedUser.Password), []byte(loginForm.Password))

	if errCompare != nil {

		context.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid Email or Password",
		})
		return
	}

	token, err := services.GenerateJWTToken(loginForm)

	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return

	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}

func HandleClaimsChecker(context *gin.Context) {

	mapClaims, exists := context.Get("mapClaims")

	if !exists {

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "mapClaims not exists",
		})
		return

	}

	//email := mapClaims["email"]

	context.JSON(http.StatusOK, gin.H{
		"claims": mapClaims,
	})

}

func HandleGoogleAuth(context *gin.Context) {

	clientId := os.Getenv("GOOGLE_AUTH_CLIENT_ID")
	redirectUri := os.Getenv("GOOGLE_REDIRECT_URI")

	url := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=profile email", clientId, redirectUri)

	fmt.Println(url)

	context.Redirect(http.StatusMovedPermanently, url)

}

func HandleGoogleAuthCallback(context *gin.Context) {

	code := context.Query("code")

	googleToken, err := services.ExchangeCodeForGoogleToken(code)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	profile, err := services.FetchGoogleUserProfile(googleToken.AccessToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
		return
	}

	/////

	checkedUser, _ := services.GetUser(types.ByEmail, profile.Email)

	if checkedUser == nil || checkedUser.ID.String() == "" {
		newUser := domain.User{
			UserName:       profile.Name,
			Email:          profile.Email,
			EmailConfirmed: true,
		}

		errCreateUser := services.CreateUser(&newUser)
		if errCreateUser != nil {
			context.JSON(http.StatusAccepted, gin.H{
				"message": errCreateUser.Error(),
			})
			return
		}
	}

	jwtToken, err := services.GenerateJWTToken(dto.UserLogin{
		Email:      profile.Email,
		RememberMe: true,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   jwtToken,
	})

}
