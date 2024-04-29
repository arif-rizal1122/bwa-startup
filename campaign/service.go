package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaings(userID int) ([]Campaign, error)
	GetCampaingByID(input GetCampaingDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaingDetailInput, inputData CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
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

func (s *service) GetCampaingByID(input GetCampaingDetailInput) (Campaign, error) {

	campaign, err := s.repository.FindByCampaignID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID
	// pembuatan slug nanti
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)
 

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}




func (s *service) 	UpdateCampaign(inputID GetCampaingDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByCampaignID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("campaign id not found")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount


	updatedCampaign, err := s.repository.Update(campaign)
    if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil

}


