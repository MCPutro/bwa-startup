package transaction

import (
	"bwa-startup/internal/entity"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/campaign"
	"bwa-startup/internal/repository/transaction"
	"context"
	"errors"
)

type transactionServiceImpl struct {
	transaction transaction.Repository
	campaign    campaign.Repository
}

func (t *transactionServiceImpl) FindByCampaignId(ctx context.Context, userId int, campId int) ([]*response.CampaignTrx, error) {
	//check owner campaign
	_, err := t.campaign.FindById(ctx, userId, campId)
	if err != nil {
		return nil, errors.New("campaign not found")
	}

	trxList, err := t.transaction.GetByCampaignId(ctx, campId)
	if err != nil {
		return nil, err
	}

	return trxList.ToCampaignTrxList(), nil
}

func (t *transactionServiceImpl) FindByUserId(ctx context.Context, userId int) ([]*entity.Transaction, error) {
	trxList, err := t.transaction.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return trxList, nil
}

func NewService(trx transaction.Repository, campaign campaign.Repository) Service {
	return &transactionServiceImpl{
		transaction: trx,
		campaign:    campaign,
	}
}
