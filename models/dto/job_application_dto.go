package dto

import (
	"JobBuddy/types"
	"time"
)

type JobApplicationForm struct {
	CompanyName       string                  `json:"company_name" binding:"required"`
	Position          string                  `json:"position" binding:"required"`
	Description       string                  `json:"description"`
	ApplicationStatus types.ApplicationStatus `json:"application_status" binding:"required"`
	ApplicationDate   time.Time               `json:"application_date" binding:"required"`
	ResponseDate      *time.Time              `json:"response_date"`
	Notes             string                  `json:"notes" binding:"required"`
}
