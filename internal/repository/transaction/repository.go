package transaction

import (
	"bwa-startup/internal/entity"
	"context"
)

type Repository interface {
	GetByCampaignId(ctx context.Context, campId int) ([]*entity.Transaction, error)
	GetByUserId(ctx context.Context, userId int) ([]*entity.Transaction, error)
}
