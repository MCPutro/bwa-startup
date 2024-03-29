package repository

import (
	"bwa-startup/config"
	"bwa-startup/internal/repository/auth"
	"bwa-startup/internal/repository/campaign"
	"bwa-startup/internal/repository/firebase"
	"bwa-startup/internal/repository/transaction"
	"bwa-startup/internal/repository/user"
	"sync"

	"gorm.io/gorm"
)

var (
	userRepository     user.Repository
	userRepositoryOnce sync.Once

	firebaseRepository     firebase.Repository
	firebaseRepositoryOnce sync.Once

	authRepository     auth.Repository
	authRepositoryOnce sync.Once

	campaignRepository     campaign.Repository
	campaignRepositoryOnce sync.Once

	transactionRepository     transaction.Repository
	transactionRepositoryOnce sync.Once
)

type Repository interface {
	UserRepository() user.Repository
	FirebaseRepository() firebase.Repository
	AuthRepository() auth.Repository
	CampaignRepository() campaign.Repository
	TransactionRepository() transaction.Repository
}

type repoManagerImpl struct {
	db  *gorm.DB
	cfg config.Config
}

func (r *repoManagerImpl) UserRepository() user.Repository {
	userRepositoryOnce.Do(func() {
		userRepository = user.NewRepository(r.db)
	})
	return userRepository
}

func (r *repoManagerImpl) FirebaseRepository() firebase.Repository {
	firebaseRepositoryOnce.Do(func() {
		firebaseRepository = firebase.NewRepository(r.cfg)
	})
	return firebaseRepository
}

func (r *repoManagerImpl) AuthRepository() auth.Repository {
	authRepositoryOnce.Do(func() {
		authRepository = auth.NewAuth(r.cfg.AuthConf())
	})
	return authRepository
}

func (r *repoManagerImpl) CampaignRepository() campaign.Repository {
	campaignRepositoryOnce.Do(func() {
		campaignRepository = campaign.NewRepository(r.db)
	})
	return campaignRepository
}

func (r *repoManagerImpl) TransactionRepository() transaction.Repository {
	transactionRepositoryOnce.Do(func() {
		transactionRepository = transaction.NewRepository(r.db)
	})
	return transactionRepository
}

func NewRepoManagerImpl(cfg config.Config, db *gorm.DB) Repository {
	return &repoManagerImpl{
		db:  db,
		cfg: cfg,
	}
}
