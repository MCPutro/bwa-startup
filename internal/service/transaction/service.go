package transaction

import (
	"bwa-startup/internal/entity"
	"context"
)

type Service interface {
	FindByCampaignId(ctx context.Context, campId int) ([]*entity.Transaction, error)
	FindByUserId(ctx context.Context, userId int) ([]*entity.Transaction, error)
}
