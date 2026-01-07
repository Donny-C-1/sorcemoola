package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name  string    `json:"name" gorm:"not null"`
	Story string    `json:"story" gorm:"not null"`
}

func (c *Campaign) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	return nil
}
