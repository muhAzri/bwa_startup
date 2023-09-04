package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(page, pageSize int) ([]Campaign, error)
	FindByUserId(userId string, page, pageSize int) ([]Campaign, error)
	FindByID(ID string) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAll(page, pageSize int) ([]Campaign, error) {
	var campaigns []Campaign
	offset := (page - 1) * pageSize

	err := r.db.Offset(offset).Preload("CampaignImages").Limit(pageSize).Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserId(userId string, page, pageSize int) ([]Campaign, error) {
	var campaigns []Campaign
	offset := (page - 1) * pageSize

	err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Offset(offset).Limit(pageSize).Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID string) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("id = ?", ID).Preload("CampaignImages").Preload("User").Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
