package transaction

import (
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"context"
)

type Service interface {
	FindByCampaignId(ctx context.Context, userId int, campId int) ([]*response.CampaignTrx, error)
	FindByUserId(ctx context.Context, userId int) ([]*response.UserTrx, error)
	Save(ctx context.Context, trx *request.Transaction) (*response.UserTrx, error)
}
