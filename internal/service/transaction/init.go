package transaction

import (
	"bwa-startup/internal/entity"
	"bwa-startup/internal/repository/transaction"
	"context"
)

type transactionServiceImpl struct {
	trx transaction.Repository
}

func (t *transactionServiceImpl) FindByCampaignId(ctx context.Context, campId int) ([]*entity.Transaction, error) {

	trxList, err := t.trx.GetByCampaignId(ctx, campId)
	if err != nil {
		return nil, err
	}

	return trxList, nil
}

func (t *transactionServiceImpl) FindByUserId(ctx context.Context, userId int) ([]*entity.Transaction, error) {
	trxList, err := t.trx.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return trxList, nil
}

func NewService(trx transaction.Repository) Service {
	return &transactionServiceImpl{
		trx: trx,
	}
}
