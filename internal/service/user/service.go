package user

import (
	"bwa-startup/internal/entity"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"context"
)

type Service interface {
	Register(ctx context.Context, input *request.RegisterUser) (*response.User, error)
	GetAll(ctx context.Context) (*[]entity.User, error)
	Login(ctx context.Context, input *request.UserLogin) (*response.User, error)
	IsEmailAvailable(ctx context.Context, email string) (bool, error)
}
