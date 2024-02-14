package service

import (
	"bwa-startup/config"
	"bwa-startup/internal/repository"
	"bwa-startup/internal/service/firebase"
	"bwa-startup/internal/service/user"

	"sync"
)

var (
	userService     user.Service
	userServiceOnce sync.Once

	firebaseService     firebase.Service
	firebaseServiceOnce sync.Once
)

type Service interface {
	UserService() user.Service
	FirebaseService() firebase.Service
}

type serviceManagerImpl struct {
	conf config.Config
	repo repository.Repository
}

// FirebaseService implements Service.
func (s *serviceManagerImpl) FirebaseService() firebase.Service {
	firebaseServiceOnce.Do(func() {
		firebaseService = firebase.NewService(s.conf, s.repo.UserRepository())
	})
	return firebaseService
}

// UserService implements Service.
func (s *serviceManagerImpl) UserService() user.Service {
	userServiceOnce.Do(func() {
		userService = user.NewService(s.repo.UserRepository(), s.conf, s.repo.AuthRepository())
	})
	return userService
}

func NewServiceManagerImpl(config config.Config, repo repository.Repository) Service {
	return &serviceManagerImpl{
		conf: config,
		repo: repo,
	}
}
