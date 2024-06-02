package middlewares

import (
	"JobBuddy/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticator() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		// Check for presence of Authorization header
		if tokenString == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			context.Abort()
			return
		}

		// Extract token from Authorization header
		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			context.Abort()
			return
		}

		token := parts[1]

		mapClaims, ok, errValidate := services.ValidateJWTToken(token)

		if errValidate != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": errValidate.Error()})
			context.Abort()
			return
		}

		if ok {

			context.Set("mapClaims", mapClaims)
			context.Next()

		} else {

			context.JSON(http.StatusInternalServerError, gin.H{"error": "JWT validation err occuurs!"})
			context.Abort()
			return

		}

		// Continue processing the request
		context.Next()
	}
}
