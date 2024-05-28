package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserName               string    `gorm:"type:varchar(100)"`
	FirstName              string    `gorm:"type:varchar(100)"`
	LastName               string    `gorm:"type:varchar(100)"`
	Email                  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	EmailConfirmationToken string
	Password               string
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
