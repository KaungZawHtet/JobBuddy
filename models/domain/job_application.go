package domain

import (
	"JobBuddy/types"

	"github.com/google/uuid"

	"time"
)

// gorm.Model definition
type JobApplication struct {
	ID                uuid.UUID               `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserEmail         string                  `gorm:"type:varchar(255);not null" json:"user_email"`
	CompanyName       string                  `gorm:"type:varchar(255);not null"`
	Position          string                  `gorm:"type:varchar(255);not null"`
	ApplicationStatus types.ApplicationStatus `gorm:"type:varchar(50);not null"`
	ApplicationDate   time.Time               `gorm:"type:date;not null"`
	ResponseDate      *time.Time              `gorm:"type:date"`
	Notes             string                  `gorm:"type:text"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
