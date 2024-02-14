package repository

import (
	"bwa-startup/config"
	"bwa-startup/internal/repository/auth"
	"bwa-startup/internal/repository/campaign"
	"bwa-startup/internal/repository/user"
	"sync"

	"gorm.io/gorm"
)

var (
	userRepository     user.Repository
	userRepositoryOnce sync.Once

	authRepository     auth.Repository
	authRepositoryOnce sync.Once

	campaignRepository     campaign.Repository
	campaignRepositoryOnce sync.Once
)

type Repository interface {
	UserRepository() user.Repository
	AuthRepository() auth.Repository
	CampaignRepository() campaign.Repository
}

type repoManagerImpl struct {
	db  *gorm.DB
	cfg config.Config
}

// UserRepository implements Repository.
func (r *repoManagerImpl) UserRepository() user.Repository {
	userRepositoryOnce.Do(func() {
		userRepository = user.NewRepository(r.db)
	})
	return userRepository
}

// AuthRepository implements Repository.
func (r *repoManagerImpl) AuthRepository() auth.Repository {
	authRepositoryOnce.Do(func() {
		authRepository = auth.NewAuth(r.cfg.AuthConf())
	})
	return authRepository
}

// CampaignRepository implements Repository.
func (r *repoManagerImpl) CampaignRepository() campaign.Repository {
	campaignRepositoryOnce.Do(func() {
		campaignRepository = campaign.NewRepository(r.db)
	})
	return campaignRepository
}

func NewRepoManagerImpl(cfg config.Config, db *gorm.DB) Repository {
	return &repoManagerImpl{
		db:  db,
		cfg: cfg,
	}
}
