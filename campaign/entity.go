package campaign

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	ID               uuid.UUID `gorm:"type:char(36);primary_key"`
	UserID           uuid.UUID `gorm:"type:char(36);index;ForeignKey:UserID"`
	Name             string    `gorm:"type:varchar(255)"`
	ShortDescription string    `gorm:"type:text"`
	Description      string    `gorm:"type:text"`
	Perks            string    `gorm:"type:text"`
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string    `gorm:"type:varchar(255);unique_index"`
	CreatedAt        time.Time `gorm:"type:datetime"`
	UpdatedAt        time.Time `gorm:"type:datetime"`
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	CampaignID uuid.UUID `gorm:"type:char(36);index;ForeignKey:CampaignID"`
	Campaign   Campaign  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time `gorm:"type:datetime"`
	UpdatedAt  time.Time `gorm:"type:datetime"`
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Campaign{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&CampaignImage{})

	if err != nil {
		return err
	}

	return nil
}
