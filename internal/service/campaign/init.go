package campaign

import (
	"bwa-startup/config"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/campaign"
	"context"
)

type campaignServiceImpl struct {
	repo campaign.Repository
	cfg  config.FirebaseConfig
}

func (s *campaignServiceImpl) GetCampaignByUserId(ctx context.Context, userId int) ([]*response.Campaign, error) {
	//check user id
	if userId <= 0 {
		return nil, nil
	}

	campaignByUserId, err := s.repo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return campaignByUserId.ToCampaignRespList(), nil
}

func (s *campaignServiceImpl) GetCampaignDetailById(ctx context.Context, userId, campaignId int) (*response.CampaignDetail, error) {
	if userId <= 0 {
		return nil, nil
	}

	campaignById, err := s.repo.FindById(ctx, userId, campaignId)
	if err != nil {
		return nil, err
	}
	
	return campaignById.ToCampaignDetailResp(s.cfg.BucketName()), nil
}

func NewService(cfg config.FirebaseConfig, repo campaign.Repository) Service {
	return &campaignServiceImpl{
		repo: repo,
		cfg:  cfg,
	}
}
