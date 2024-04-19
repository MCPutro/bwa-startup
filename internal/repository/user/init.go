package user

import (
	newError "bwa-startup/internal/common/errors"
	"bwa-startup/internal/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func (ur *repositoryImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := ur.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *repositoryImpl) FindById(ctx context.Context, ID int) (*entity.User, error) {
	var user entity.User

	result := ur.db.WithContext(ctx).First(&user, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, newError.ErrUserIdNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *repositoryImpl) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User

	result := ur.db.WithContext(ctx).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		return users, nil
	}
	// fmt.Println("Found ", result.RowsAffected, "(s) data")

	return nil, nil
}

func (ur *repositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	result := ur.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, newError.ErrEmailNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *repositoryImpl) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := ur.db.WithContext(ctx).Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
