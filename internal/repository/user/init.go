package user

import (
	"bwa-startup/internal/entity"
	"context"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

// Save implements UserRepository.
func (ur *repositoryImpl) Save(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := ur.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindById implements UserRepository.
func (ur *repositoryImpl) FindById(ctx context.Context, ID int) (*entity.User, error) {
	var user entity.User

	err := ur.db.WithContext(ctx).First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll implements UserRepository.
func (ur *repositoryImpl) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User

	result := ur.db.WithContext(ctx).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	// fmt.Println("Found ", result.RowsAffected, "(s) data")

	return users, nil
}

// FindByEmail implements Repository.
func (ur *repositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := ur.db.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
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
