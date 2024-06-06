package services

import (
	"JobBuddy/config"
	"JobBuddy/models/domain"
	"JobBuddy/models/dto"
	"JobBuddy/types"
	"time"
)

func ListAllMyJobApplication(email string, search string, status types.ApplicationStatus, startDate time.Time, endDate time.Time, limit int, offset int) ([]domain.JobApplication, error) {
	var jobApps []domain.JobApplication
	db, errDbAccess := config.AcessDB()
	if errDbAccess != nil {
		return nil, errDbAccess
	}

	query := db.Where("user_email = ?", email)

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

func CreateMyApplicationForm(email string, dto dto.JobApplicationForm) error {

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
