package handlers

import (
	"JobBuddy/services"
	"JobBuddy/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"time"
)

func HandleMyApplicationsList(context *gin.Context) {

	// Get the search query parameter
	search := context.Query("search")

	// Get the status query parameter
	status := context.Query("status")
	var applicationStatus types.ApplicationStatus
	if status != "" {
		applicationStatus = types.ApplicationStatus(status)
	}

	// Get the start date query parameter
	startDateStr := context.Query("start_date")
	var startDate time.Time
	if startDateStr != "" {
		startDate, _ = time.Parse("2006-01-02", startDateStr)
	}

	// Get the end date query parameter
	endDateStr := context.Query("end_date")
	var endDate time.Time
	if endDateStr != "" {
		endDate, _ = time.Parse("2006-01-02", endDateStr)
	}

	// Get the limit query parameter
	limit, err := strconv.Atoi(context.Query("limit"))
	if err != nil {
		limit = 10
	}

	// Get the offset query parameter
	offset, err := strconv.Atoi(context.Query("offset"))
	if err != nil {
		offset = 0
	}

	claims, exists := context.Get("mapClaims")

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorization detected",
		})
	}

	mapClaims, ok := claims.(jwt.MapClaims)

	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cliam Error",
		})
	}

	email := mapClaims["email"].(string)

	// Call the GetAllMyJobApplication service
	applications, err := services.GetAllMyJobApplication(email, search, applicationStatus, startDate, endDate, limit, offset)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of applications in JSON format
	context.JSON(http.StatusOK, gin.H{"message": "This is your applications", "data": applications})

}
