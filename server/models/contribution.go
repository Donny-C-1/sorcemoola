package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contribution struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Amount     int64     `json:"amount" gorm:"not null"`
	UserID     uuid.UUID `json:"userId" gorm:"not null;index"`
	CampaignID uuid.UUID `json:"campaignId" gorm:"not null;index"`
	CreatedAt  time.Time `json:"createdAt"`

	User     User     `gorm:"foreignKey:UserID"`
	Campaign Campaign `gorm:"foreignKey:CampaignID"`
}

func (c *Contribution) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	return nil
}
