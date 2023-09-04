package campaign

type Service interface {
	GetCampaign(userID *string, page *int, limit *int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
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
