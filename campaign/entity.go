package campaign

import (
	"time"

	"github.com/arif-rizal1122/bwa-startup/user"
)

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt 		 time.Time
	CampaignImages   []CampaignImage
	User 			 user.User
}



type CampaignImage struct {
	// fields here
	ID 				int
	CampaignID      int
	FileName		string
	IsPrimary 		int
	CreatedAt        time.Time
	UpdatedAt 		 time.Time
}