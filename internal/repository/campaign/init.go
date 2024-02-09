package campaign

import (
	"bwa-startup/internal/entity"
	"context"
	"gorm.io/gorm"
)

type campaignImpl struct {
	db *gorm.DB
}

func (c *campaignImpl) FindAll(ctx context.Context) (*[]entity.Campaign, error) {
	var campaigns []entity.Campaign
	err := c.db.WithContext(ctx).Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return &campaigns, nil
}

func (c *campaignImpl) FindByUserId(ctx context.Context, userId int) (*[]entity.Campaign, error) {
	var campaigns []entity.Campaign
	err := c.db.WithContext(ctx).Where("user_id = ?", userId).Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return &campaigns, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &campaignImpl{db: db}
}
