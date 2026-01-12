package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name         string    `json:"name" gorm:"not null"`
	Description  string    `json:"description"`
	Story        string    `json:"story" gorm:"not null"`
	Slug         string    `json:"slug" gorm:"not null;uniquIndex"`
	FundGoal     int64     `json:"fundGoal" gorm:"not null"`
	AmountRaised int64     `json:"amountRaised" gorm:"default:0"`
	BackersCount int       `json:"backersCount" gorm:"default:0"`
	CreatedBy    string    `json:"createdBy" gorm:"not null;index"`
	Status       string    `json:"status" gorm:"default:'active'"`

	// Relationship
	Creator             User           `json:"creator" gorm:"foreignKey:CreatedBy"`
	RecentContributions []Contribution `json:"recentContributions" gorm:"foreignKey:CampaignID"`
}

func (c *Campaign) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	c.Slug = c.ID.String()

	return nil
}
