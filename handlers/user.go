package handlers

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	//"JobBuddy/config"
	"JobBuddy/models/domain"
	"JobBuddy/models/dto"
	"JobBuddy/services"
	"JobBuddy/types"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
	}

}
