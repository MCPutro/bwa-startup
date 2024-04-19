package request

import "bwa-startup/internal/entity"

type Transaction struct {
	Amount     int `json:"amount"`
	CampaignId int `json:"campaign_id"`
	UserId     int `json:"-"`
}

func (t *Transaction) ToEntity() *entity.Transaction {
	return &entity.Transaction{
		CampaignId: t.CampaignId,
		Amount:     t.Amount,
		UserId:     t.UserId,
		Status:     "Unpaid",
	}
}
