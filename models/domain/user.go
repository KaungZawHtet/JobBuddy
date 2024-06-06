package domain

import (
	"time"

	"github.com/google/uuid"
)

// gorm.Model definition
type User struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserName               string    `gorm:"type:varchar(100)"`
	FirstName              string    `gorm:"type:varchar(100)"`
	LastName               string    `gorm:"type:varchar(100)"`
	Email                  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	EmailConfirmationToken string
	EmailConfirmed         bool
	Password               string
	CreatedAt              time.Time
	UpdatedAt              time.Time

	JobApplications []JobApplication `gorm:"foreignKey:UserID"`
}
