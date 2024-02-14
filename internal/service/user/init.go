package user

import (
	"bwa-startup/config"
	"bwa-startup/internal/entity"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/auth"
	"bwa-startup/internal/repository/user"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type serviceImpl struct {
	repository user.Repository
	config     config.Config
	auth       auth.Repository
}

// GetAll implements Service.
func (us *serviceImpl) GetAll(ctx context.Context) ([]*entity.User, error) {
	// resp = new([]entity.User)

	u, err := us.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Register implements UserService.
func (us *serviceImpl) Register(ctx context.Context, input *request.RegisterUser) (*response.User, error) {
	//check email is already register
	isExistingUser, err := us.IsEmailAvailable(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if !isExistingUser {
		return nil, errors.New("email already register")
	}

	//save new user to database
	userEntity := input.ToEntity()
	u, err := us.repository.Save(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	//generate token
	token, err := us.auth.GenerateToken(u)
	if err != nil {
		return nil, err
	}

	return u.ToUserResponse(us.config.FirebaseConf().BucketName(), token), nil
}

// Login implements Service.
func (us *serviceImpl) Login(ctx context.Context, input *request.UserLogin) (*response.User, error) {
	existingUser, err := us.repository.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	// if user is found, then check password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("email and password not match")
	}

	//generate token
	token, err := us.auth.GenerateToken(existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser.ToUserResponse(us.config.FirebaseConf().BucketName(), token), nil
}

// IsEmailAvailable implements Service.
func (us *serviceImpl) IsEmailAvailable(ctx context.Context, email string) (bool, error) {
	//get user by email
	existingUser, err := us.repository.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if existingUser != nil && existingUser.ID > 0 {
		return false, nil
	}

	return true, nil
}

func NewService(repo user.Repository, conf config.Config, auth auth.Repository) Service {
	return &serviceImpl{
		repository: repo,
		config:     conf,
		auth:       auth,
	}
}
