package transaction

import (
	"bwa-startup/internal/entity"
	"context"

	"gorm.io/gorm"
)

type transactionImpl struct {
	db *gorm.DB
}

func (t *transactionImpl) GetByCampaignId(ctx context.Context, campId int) (entity.TransactionList, error) {
	var tmp []*entity.Transaction

	result := t.db.WithContext(ctx).Where("campaign_id = ?", campId).Preload("User").Order("created_at desc").Find(&tmp)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return tmp, nil
	}
	return nil, nil
}

func (t *transactionImpl) GetByUserId(ctx context.Context, userId int) (entity.TransactionList, error) {
	var tmp []*entity.Transaction

	result := t.db.WithContext(ctx).Where("user_id = ?", userId).Preload("Campaign").Order("created_at desc").Find(&tmp)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return tmp, nil
	}
	return nil, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &transactionImpl{db: db}
}
