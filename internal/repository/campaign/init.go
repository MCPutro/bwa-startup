package campaign

import (
	newError "bwa-startup/internal/common/errors"
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

	result := c.db.WithContext(ctx).Where("user_id = ?", userId).Preload("CampaignImages").Find(&campaigns)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		return campaigns, nil
	}

	return nil, nil
}

func (c *campaignImpl) FindById(ctx context.Context, userId, campaignId int) (*entity.Campaign, error) {
	campaign := entity.Campaign{}

	result := c.db.WithContext(ctx).Where("id = ? and user_id = ?", campaignId, userId).Preload("User").Preload("CampaignImages").Find(&campaign)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		return &campaign, nil
	}

	return nil, newError.ErrNotFound

}

func (c *campaignImpl) Save(ctx context.Context, campaign *entity.Campaign) (*entity.Campaign, error) {
	err := c.db.WithContext(ctx).Create(campaign).Error
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (c *campaignImpl) Update(ctx context.Context, campaign *entity.Campaign) (*entity.Campaign, error) {
	err := c.db.WithContext(ctx).Save(campaign).Error
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (c *campaignImpl) CreateImage(ctx context.Context, image *entity.CampaignImage) error {
	//check new image is primary or not
	if image.IsPrimary {
		tx := c.db.WithContext(ctx).Model(&entity.CampaignImage{}).Where("campaign_id", image.CampaignID).Update("is_primary", false)
		if tx.Error != nil {
			return tx.Error
		}
	}

	err := c.db.WithContext(ctx).Create(&image).Error
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *gorm.DB) Repository {
	return &campaignImpl{db: db}
}
