package campaign

import (
	"bwa-startup/internal/entity"
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) (entity.CampaignList, error)
	FindByUserId(ctx context.Context, userId int) (entity.CampaignList, error)
}
