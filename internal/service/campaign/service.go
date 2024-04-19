package campaign

import (
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"context"
	"mime/multipart"
)

type Service interface {
	GetByUserId(ctx context.Context, userId int) ([]*response.Campaign, error)
	GetDetailById(ctx context.Context, userId, campaignId int) (*response.CampaignDetail, error)
	Save(ctx context.Context, campaign *request.Campaign) (*response.CampaignDetail, error)
	Update(ctx context.Context, campaignId int, newCampaign *request.Campaign) (*response.CampaignDetail, error)
	UploadImage(ctx context.Context, userId, campaignId int, file *multipart.FileHeader, isPrimary bool) error
}
