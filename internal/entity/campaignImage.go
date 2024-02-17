package entity

import "time"

type CampaignImage struct {
	ID         int
	CampaignID int
	Filename   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (i *CampaignImage) TableName() string {
	return "campaign_images"
}
