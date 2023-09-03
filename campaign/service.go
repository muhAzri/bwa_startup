package campaign

type Service interface {
	FindCampaign(userID *string, page *int, limit *int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaign(userID *string, page *int, limit *int) ([]Campaign, error) {
	if page == nil {
		defaultPage := 1
		page = &defaultPage
	}

	if limit == nil {
		defaultLimit := 10
		limit = &defaultLimit
	}

	if userID != nil {
		return s.repository.FindByUserId(*userID, *page, *limit)
	}

	return s.repository.FindAll(*page, *limit)
}
