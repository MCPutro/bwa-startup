package entity

import (
	"bwa-startup/internal/constants"
	"bwa-startup/internal/handler/response"
	"time"
)

type Transaction struct {
	ID         int `gorm:"primarykey"`
	CampaignId int
	Campaign   Campaign
	UserId     int
	User       User
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}

type TransactionList []*Transaction

func (tl *TransactionList) ToCampaignTrxList() []*response.CampaignTrx {
	var resp []*response.CampaignTrx
	for _, v := range *tl {
		resp = append(resp, &response.CampaignTrx{
			Id:        v.ID,
			Name:      v.User.Name,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Format(constants.DatetimeFormat),
		})
	}
	return resp
}

func (tl *TransactionList) ToUserTrxList() []*response.UserTrx {
	var resp []*response.UserTrx
	for _, v := range *tl {
		resp = append(resp, &response.UserTrx{
			Id:           v.ID,
			CampaignName: v.Campaign.Name,
			Amount:       v.Amount,
			Status:       v.Status,
			CreatedAt:    v.CreatedAt.Format(constants.DatetimeFormat),
		})
	}
	return resp
}
