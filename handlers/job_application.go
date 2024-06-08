package handlers

import (
	"JobBuddy/models/dto"
	"JobBuddy/services"
	"JobBuddy/types"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

	userEmail, errClaim := services.GetClaimVar(types.ByEmail, context)

	if errClaim != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errClaim.Error()})
		return
	}

	strUserEmail := userEmail.(string)

	// Call the ListAllMyJobApplication service
	applications, err := services.ListAllMyJobApplication(strUserEmail, search, applicationStatus, startDate, endDate, limit, offset)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the list of applications in JSON format
	context.JSON(http.StatusOK, gin.H{"message": "This is your applications", "data": applications})

}

func HandleMyApplicationCreation(context *gin.Context) {

	var jobAppForm dto.JobApplicationForm

	errBindJson := context.ShouldBindJSON(&jobAppForm)

	if errBindJson != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errBindJson.Error()})
		return
	}

	userEmail, errClaim := services.GetClaimVar(types.ByEmail, context)

	if errClaim != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errClaim.Error()})
		return
	}

	strUserEmail := userEmail.(string)

	err := services.CreateMyJobApplicationForm(strUserEmail, jobAppForm)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully created your Job Application",
		"data":    jobAppForm,
	})

}

func HandleMyJobApplicationDeletion(context *gin.Context) {

	id := context.Param("id")

	userEmail, errClaim := services.GetClaimVar(types.ByEmail, context)

	if errClaim != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errClaim.Error()})
		return
	}

	strUserEmail := userEmail.(string)

	errDelete := services.DeleteMyJobApplication(strUserEmail, id)

	if errDelete != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errDelete.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted your Job Application",
	})

}

func HandleMyJobApplicationPatch(context *gin.Context) {

	id := context.Param("id")

	var jobAppForm dto.JobApplicationForm

	errBindJson := context.ShouldBindJSON(&jobAppForm)

	if errBindJson != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errBindJson.Error()})
		return
	}

	userEmail, errClaim := services.GetClaimVar(types.ByEmail, context)

	if errClaim != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": errClaim.Error()})
		return
	}

	strUserEmail := userEmail.(string)

	errDelete := services.PatchMyJobApplication(strUserEmail, id, jobAppForm)

	if errDelete != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": errDelete.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted your Job Application",
	})

}
