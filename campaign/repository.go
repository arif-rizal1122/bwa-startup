package campaign

import (

	"gorm.io/gorm"
)

type Repository interface {
	// []campaign, mengembalikan lebih dari satu data camapign di db 
	FindAll() ([]Campaign, error)
	FindByUserID(ID int) ([]Campaign, error)
	FindByCampaignID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
}

type repository struct {
	// fields here
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}






func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, err
}





func (r *repository) FindByUserID(ID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", ID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}



func (r repository) FindByCampaignID(ID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}



func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

