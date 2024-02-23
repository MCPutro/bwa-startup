package campaign

import (
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"context"
	"mime/multipart"
)

type Service interface {
	GetCampaignByUserId(ctx context.Context, userId int) ([]*response.Campaign, error)
	GetCampaignDetailById(ctx context.Context, userId, campaignId int) (*response.CampaignDetail, error)
	CreateCampaign(ctx context.Context, campaign *request.Campaign) (*response.CampaignDetail, error)
	UpdateCampaign(ctx context.Context, campaignId int, newCampaign *request.Campaign) (*response.CampaignDetail, error)
	UploadImage(ctx context.Context, userId, campaignId int, file multipart.File, fileHeader *multipart.FileHeader, isPrimary bool) error
}
