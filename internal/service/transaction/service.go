package transaction

import (
	"bwa-startup/internal/entity"
	"bwa-startup/internal/handler/response"
	"context"
)

type Service interface {
	FindByCampaignId(ctx context.Context, userId int, campId int) ([]*response.CampaignTrx, error)
	FindByUserId(ctx context.Context, userId int) ([]*entity.Transaction, error)
}
