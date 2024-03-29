package transaction

import (
	"bwa-startup/internal/entity"
	"context"
	"gorm.io/gorm"
)

type transactionImpl struct {
	db *gorm.DB
}

func (t *transactionImpl) GetByCampaignId(ctx context.Context, campId int) ([]*entity.Transaction, error) {
	var tmp []*entity.Transaction

	result := t.db.WithContext(ctx).Where("campaign_id = ?", campId).Find(&tmp).Preload("User")
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected > 0 {
		return tmp, nil
	}
	return nil, nil
}

func (t *transactionImpl) GetByUserId(ctx context.Context, userId int) ([]*entity.Transaction, error) {
	var tmp []*entity.Transaction

	result := t.db.WithContext(ctx).Where("campaign_id = ?", userId).Find(&tmp).Preload("User")
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
