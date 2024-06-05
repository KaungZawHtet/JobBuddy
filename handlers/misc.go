package handlers

import "github.com/gin-gonic/gin"
import "net/http"

func HandlePing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{

		"message": "pong",
	})

}
