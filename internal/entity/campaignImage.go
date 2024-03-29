package entity

import "time"

type CampaignImage struct {
	ID         int
	CampaignID int
	Image      string
	Token      string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (i *CampaignImage) TableName() string {
	return "campaign_images"
}
