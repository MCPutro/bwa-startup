package service

import (
	"bwa-startup/config"
	"bwa-startup/internal/repository"
	"bwa-startup/internal/service/campaign"
	"bwa-startup/internal/service/user"

	"sync"
)

var (
	userService     user.Service
	userServiceOnce sync.Once

	campaignService     campaign.Service
	campaignServiceOnce sync.Once
)

type Service interface {
	UserService() user.Service

	CampaignService() campaign.Service
}

type serviceManagerImpl struct {
	conf config.Config
	repo repository.Repository
}

// UserService implements Service.
func (s *serviceManagerImpl) UserService() user.Service {
	userServiceOnce.Do(func() {
		userService = user.NewService(s.conf, s.repo.UserRepository(), s.repo.AuthRepository(), s.repo.FirebaseRepository())
	})
	return userService
}

func (s *serviceManagerImpl) CampaignService() campaign.Service {
	campaignServiceOnce.Do(func() {
		campaignService = campaign.NewService(s.conf.FirebaseConf(), s.repo.CampaignRepository(), s.repo.UserRepository(), s.repo.FirebaseRepository())
	})
	return campaignService
}

func NewServiceManagerImpl(config config.Config, repo repository.Repository) Service {
	return &serviceManagerImpl{
		conf: config,
		repo: repo,
	}
}
