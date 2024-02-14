package campaign

import (
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/campaign"
	"context"
)

type serviceImpl struct {
	repo campaign.Repository
}

func (s *serviceImpl) GetCampaignByUserId(ctx context.Context, userId int) ([]*response.Campaign, error) {
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

func NewService(repo campaign.Repository) Service {
	return &serviceImpl{repo: repo}
}
