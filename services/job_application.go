package services

import (
	"JobBuddy/config"
	"JobBuddy/models/domain"
	"JobBuddy/types"
	"time"
)

func GetAllMyJobApplication(email string, search string, status types.ApplicationStatus, startDate time.Time, endDate time.Time, limit int, offset int) ([]domain.JobApplication, error) {

	var jobApps []domain.JobApplication
	db, errDbAccess := config.AcessDB()

	if errDbAccess != nil {

		return nil, errDbAccess

	}

	db = db.Where("email = ?", email).Find(&jobApps)

	if limit > 0 {
		db = db.Limit(limit)
	}

	if offset > 0 {
		db = db.Offset(offset)
	}

	if len(search) > 0 {
		db = db.Where("company_name LIKE %?%", search).Find(&jobApps)
		if len(jobApps) == 0 {

			db = db.Where("position LIKE %?%", search).Find(&jobApps)

		}
	}

	if len(status) > 0 {
		db = db.Where("application_status = ?", status).Find(&jobApps)
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&jobApps)
	}

	return jobApps, nil

}
