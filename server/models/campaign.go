package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
}

func (c *Campaign) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	return nil
}
