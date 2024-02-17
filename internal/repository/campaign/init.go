package campaign

import (
	"bwa-startup/internal/entity"
	"context"
	"gorm.io/gorm"
)

type campaignImpl struct {
	db *gorm.DB
}

func (c *campaignImpl) FindAll(ctx context.Context) (entity.CampaignList, error) {
	var campaigns []*entity.Campaign
	err := c.db.WithContext(ctx).Preload("CampaignImages").Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (c *campaignImpl) FindByUserId(ctx context.Context, userId int) (entity.CampaignList, error) {
	var campaigns []*entity.Campaign

	err := c.db.WithContext(ctx).Where("user_id = ?", userId).Preload("CampaignImages").Find(&campaigns).Error
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (c *campaignImpl) FindById(ctx context.Context, userId, campaignId int) (*entity.Campaign, error) {

	campaign := entity.Campaign{}

	err := c.db.WithContext(ctx).Where("id = ? and user_id = ?", campaignId, userId).Preload("User").Preload("CampaignImages").Find(&campaign).Error
	if err != nil {
		return nil, err
	}

	return &campaign, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &campaignImpl{db: db}
}
