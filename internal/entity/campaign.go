package entity

import (
	"bwa-startup/internal/handler/response"
	"time"
)

type Campaign struct {
	ID              int
	UserId          int
	Name            string
	SortDescription string
	Description     string
	Perks           string
	BackerCount     int
	GoalAmount      int
	CurrentAmount   int
	Slug            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CampaignImages  []CampaignImage
}

func (c *Campaign) ToRespCampaign() *response.Campaign {
	url := ""
	for i := 0; i < len(c.CampaignImages); i++ {
		if c.CampaignImages[i].IsPrimary == true {
			url = c.CampaignImages[i].Filename
			break
		}
	}

	return &response.Campaign{
		Id:               c.ID,
		Title:            c.Name,
		ShortDescription: c.SortDescription,
		Description:      c.Description,
		GoalAmount:       c.GoalAmount,
		CurrentAmount:    c.CurrentAmount,
		UserId:           c.UserId,
		ImageUrl:         url,
	}
}

type CampaignList []*Campaign

func (c *CampaignList) ToCampaignRespList() []*response.Campaign {
	var temp []*response.Campaign
	for _, campaign := range *c {
		temp = append(temp, campaign.ToRespCampaign())
	}
	return temp
}
