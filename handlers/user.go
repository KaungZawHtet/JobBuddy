package handlers

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	"JobBuddy/models/dto"
	"net/http"
)

func HandleRegister(context *gin.Context) {

	var userRegistration dto.UserRegistration

	err := context.ShouldBindJSON(&userRegistration)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	context.JSON(http.StatusAccepted, gin.H{
		"message": "Check your email",
		"data":    userRegistration,
	})

}
