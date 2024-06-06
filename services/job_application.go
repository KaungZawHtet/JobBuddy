package services

import (
	"JobBuddy/config"
	"JobBuddy/models/domain"
	"JobBuddy/models/dto"
	"JobBuddy/types"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func ListAllMyJobApplication(userEmail string, search string, status types.ApplicationStatus, startDate time.Time, endDate time.Time, limit int, offset int) ([]domain.JobApplication, error) {
	var jobApps []domain.JobApplication
	db, errDbAccess := config.AcessDB()
	if errDbAccess != nil {
		return nil, errDbAccess
	}

	query := db.Where("user_email = ?", userEmail)

	if limit > 0 {
		query = query.Limit(limit)
	}

	if offset > 0 {
		query = query.Offset(offset)
	}

	if len(search) > 0 {
		searchPattern := "%" + search + "%"
		query = query.Where("company_name LIKE ? OR position LIKE ?", searchPattern, searchPattern)
	}

	if status != "" {
		query = query.Where("application_status = ?", status)
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Execute the query
	if err := query.Find(&jobApps).Error; err != nil {
		return nil, err
	}

	return jobApps, nil
}

func CreateMyJobApplicationForm(email string, dto dto.JobApplicationForm) error {

	db, errDbAccess := config.AcessDB()

	if errDbAccess != nil {

		return errDbAccess

	}

	db.Save(&domain.JobApplication{
		UserEmail:         email,
		CompanyName:       dto.CompanyName,
		Position:          dto.Position,
		ApplicationStatus: dto.ApplicationStatus,
		ApplicationDate:   dto.ApplicationDate,
		ResponseDate:      dto.ResponseDate,
		Notes:             dto.Notes,
	})

	return nil

}

func DeleteMyJobApplication(userEmail string, id string) error {
	db, errDbAccess := config.AcessDB()
	if errDbAccess != nil {
		return errDbAccess
	}

	var checkedJobApp domain.JobApplication

	// Convert string ID to UUID
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid UUID format: %v", err)
	}
	db.Where(" id = ?", uid).Find(&checkedJobApp)

	if checkedJobApp.UserEmail != userEmail {
		return fmt.Errorf("unauthorized to access the data")
	}

	// Perform the delete operation
	result := db.Delete(&domain.JobApplication{}, uid)
	if result.Error != nil {
		return result.Error
	}

	// Check if any record was deleted
	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with ID: %v", id)
	}

	return nil
}

func PatchMyJobApplication(userEmail string, id string, dto dto.JobApplicationForm) error {

	var jobApp domain.JobApplication

	db, errDbAccess := config.AcessDB()
	if errDbAccess != nil {
		return errDbAccess
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid UUID format: %v", err)
	}

	db.Where("id = ?", uid).Find(&jobApp)

	if jobApp.UserEmail != userEmail {
		return fmt.Errorf("unauthorized to access the record")
	}

	jobApp.CompanyName = dto.CompanyName
	jobApp.ApplicationDate = dto.ApplicationDate
	jobApp.ApplicationStatus = dto.ApplicationStatus
	jobApp.Notes = dto.Notes
	jobApp.Position = dto.Position
	jobApp.ResponseDate = dto.ResponseDate

	result := db.Save(&jobApp)

	if result.RowsAffected == 0 {
		return fmt.Errorf("unable to path the Job Application with ID: %v", id)
	}

	return nil

}
