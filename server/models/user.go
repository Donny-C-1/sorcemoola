package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FirstName     string         `json:"first_name" gorm:"not null"`
	LastName      string         `json:"last_name" gorm:"not null"`
	PreferredName string         `json:"preferred_name"`
	Email         string         `gorm:"unique;not null"`
	PasswordHash  string         `json:"-" gorm:"not null"`
	DateOfBirth   time.Time      `json:"date_of_birth"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
