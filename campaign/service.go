package campaign

type Service interface {
	GetCampaings(userID int) ([]Campaign, error)

}

type service struct{
	 repository Repository
}

func NewService (repository Repository) *service {
	return &service{repository: repository}
}




func (s *service) GetCampaings(userID int) ([]Campaign, error) {
	// kalau userID ada
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	// kalau tidak ada userID
	compaigns, err := s.repository.FindAll()
	
	if err != nil {
		return compaigns, err
	}
	return compaigns, nil
}