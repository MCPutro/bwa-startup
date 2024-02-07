package repository

import (
	"bwa-startup/config"
	"bwa-startup/internal/repository/jwt"
	"bwa-startup/internal/repository/user"
	"sync"

	"gorm.io/gorm"
)

var (
	userRepository     user.Repository
	userRepositoryOnce sync.Once

	JwtRepository     jwt.Repository
	JwtRepositoryOnce sync.Once
)

type Repository interface {
	UserRepository() user.Repository
	JwtRepository() jwt.Repository
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

// JwtRepository implements Repository.
func (r *repoManagerImpl) JwtRepository() jwt.Repository {
	JwtRepositoryOnce.Do(func() {
		JwtRepository = jwt.NewJWT(r.cfg.AuthConf())
	})
	return JwtRepository
}

func NewRepoManagerImpl(cfg config.Config, db *gorm.DB) Repository {
	return &repoManagerImpl{
		db:  db,
		cfg: cfg,
	}
}
