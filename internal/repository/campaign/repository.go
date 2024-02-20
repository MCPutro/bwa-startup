package campaign

import (
	"bwa-startup/internal/entity"
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) (entity.CampaignList, error)
	FindByUserId(ctx context.Context, userId int) (entity.CampaignList, error)
	FindById(ctx context.Context, userId, campaignId int) (*entity.Campaign, error)
	Save(ctx context.Context, campaign *entity.Campaign) (*entity.Campaign, error)
	Update(ctx context.Context, campaign *entity.Campaign) (*entity.Campaign, error)
	CreateImage(ctx context.Context, image *entity.CampaignImage) (entity.CampaignImage, error)
}
