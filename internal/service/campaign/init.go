package campaign

import (
	"bwa-startup/config"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/campaign"
	"bwa-startup/internal/repository/user"
	"context"
	"errors"
)

type campaignServiceImpl struct {
	campaign campaign.Repository
	user     user.Repository
	cfg      config.FirebaseConfig
}

func (s *campaignServiceImpl) GetCampaignByUserId(ctx context.Context, userId int) ([]*response.Campaign, error) {
	//check user id
	existingUser, _ := s.user.FindById(ctx, userId)
	if existingUser == nil {
		return nil, errors.New("user id not found")
	}

	campaignByUserId, err := s.campaign.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return campaignByUserId.ToCampaignRespList(), nil
}

func (s *campaignServiceImpl) GetCampaignDetailById(ctx context.Context, userId, campaignId int) (*response.CampaignDetail, error) {
	if userId <= 0 {
		return nil, nil
	}

	campaignById, err := s.campaign.FindById(ctx, userId, campaignId)
	if err != nil {
		return nil, err
	}

	return campaignById.ToCampaignDetailResp(s.cfg.BucketName()), nil
}

func (s *campaignServiceImpl) CreateCampaign(ctx context.Context, campaign *request.Campaign) (*response.CampaignDetail, error) {
	//check user
	existingUser, _ := s.user.FindById(ctx, campaign.UserId)
	if existingUser == nil {
		return nil, errors.New("user id not found")
	}

	save, err := s.campaign.Save(ctx, campaign.ToEntity())
	if err != nil {
		return nil, err
	}

	save.User = *existingUser

	return save.ToCampaignDetailResp(s.cfg.BucketName()), nil
}

func NewService(cfg config.FirebaseConfig, repo campaign.Repository, user user.Repository) Service {
	return &campaignServiceImpl{
		campaign: repo,
		user:     user,
		cfg:      cfg,
	}
}
