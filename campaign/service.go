package campaign

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaign(userID *string, page *int, limit *int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaign(userID *string, page *int, limit *int) ([]Campaign, error) {
	if page == nil {
		defaultPage := 1
		page = &defaultPage
	}

	if limit == nil {
		defaultLimit := 10
		limit = &defaultLimit
	}

	if userID != nil {
		campaigns, err := s.repository.FindByUserId(*userID, *page, *limit)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll(*page, *limit)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	slugString := fmt.Sprintf("%s %s", input.Name, input.User.ID)

	campaign := Campaign{
		ID:               uuid.New(),
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		UserID:           input.User.ID,
		Slug:             slug.Make(slugString),
	}

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
