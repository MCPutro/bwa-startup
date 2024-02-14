package user

import (
	"bwa-startup/internal/entity"
	"context"
)

type Repository interface {
	Save(ctx context.Context, user *entity.User) (*entity.User, error)
	FindById(ctx context.Context, ID int) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
}
