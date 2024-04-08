package transaction

import (
	"bwa-startup/internal/entity"
	"context"
)

type Repository interface {
	GetByCampaignId(ctx context.Context, campId int) (entity.TransactionList, error)
	GetByUserId(ctx context.Context, userId int) (entity.TransactionList, error)
}
