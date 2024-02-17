package campaign

import (
	"bwa-startup/internal/handler/response"
	"context"
)

type Service interface {
	GetCampaignByUserId(ctx context.Context, userId int) ([]*response.Campaign, error)
	GetCampaignDetailById(ctx context.Context, userId, campaignId int) (*response.CampaignDetail, error)
}
