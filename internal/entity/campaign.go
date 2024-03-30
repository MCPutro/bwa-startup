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

func (c *Campaign) ToCampaignDetailResp() *response.CampaignDetail {
	imagePrimary := ""
	var listImage []string
	for i := 0; i < len(c.CampaignImages); i++ {
		listImage = append(listImage, c.CampaignImages[i].Image) //common.GetUrlImage(bucket, c.CampaignImages[i].Image, c.CampaignImages[i].Token))
		if c.CampaignImages[i].IsPrimary && imagePrimary == "" {
			imagePrimary = c.CampaignImages[i].Image //common.GetUrlImage(bucket, c.CampaignImages[i].Image, c.CampaignImages[i].Token)
			//break
			//continue
		}
	}

	var username, avatar string
	var userId int
	if c.User != (User{}) {
		userId = c.User.ID
		username = c.User.Name
		avatar = c.User.Image
	}

	return &response.CampaignDetail{
		Id:               c.ID,
		Title:            c.Name,
		ShortDescription: c.SortDescription,
		Description:      c.Description,
		ImageUrl:         imagePrimary,
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
	imageUrl := ""
	for i := 0; i < len(c.CampaignImages); i++ {
		if c.CampaignImages[i].IsPrimary {
			imageUrl = c.CampaignImages[i].Image //common.GetUrlImage(bucket, c.CampaignImages[i].Image, c.CampaignImages[i].Token)
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
		ImageUrl:         imageUrl,
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
