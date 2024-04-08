package response

type CampaignTrx struct {
	Id        int    `json:"transaction_id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}
