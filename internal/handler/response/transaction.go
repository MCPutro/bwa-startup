package response

import "bwa-startup/internal/constants"

type CampaignTrx struct {
	Id        int    `json:"transaction_id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

type UserTrx struct {
	Id           int                 `json:"transaction_id"`
	CampaignName string              `json:"campaign_name"`
	Amount       int                 `json:"amount"`
	Status       constants.TrxStatus `json:"status"`
	CreatedAt    string              `json:"created_at"`
}
