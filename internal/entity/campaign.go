package entity

import (
	"bwa-startup/internal/handler/response"
	"strings"
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
	User            User
}

func (c *Campaign) TableName() string {
	return "campaigns"
}

func (c *Campaign) ToCampaignDetailResp(bucket string) *response.CampaignDetail {
	url := ""
	var listImage []string
	for i := 0; i < len(c.CampaignImages); i++ {
		listImage = append(listImage, c.CampaignImages[i].Filename)
		if c.CampaignImages[i].IsPrimary == true && url == "" {
			url = c.CampaignImages[i].Filename
			//break
			//continue
		}
	}

	var username, avatar string
	var userId int
	if c.User != (User{}) {
		userId = c.User.ID
		username = c.User.Name
		avatar = c.User.GetUrlAvatar(bucket)
	}

	return &response.CampaignDetail{
		Id:               c.ID,
		Title:            c.Name,
		ShortDescription: c.SortDescription,
		Description:      c.Description,
		ImageUrl:         url,
		GoalAmount:       c.GoalAmount,
		CurrentAmount:    c.CurrentAmount,
		UserId:           userId,   //c.User.ID,
		UserName:         username, //c.User.Name,
		UserAvatar:       avatar,   //c.User.GetUrlAvatar(bucket),
		Perks:            strings.Split(c.Perks, "|"),
		Image:            listImage,
	}
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
