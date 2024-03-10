package user

import (
	"bwa-startup/internal/entity"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"context"
	"mime/multipart"
)

type Service interface {
	Register(ctx context.Context, newUser *request.RegisterUser) (*response.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	Login(ctx context.Context, newUser *request.UserLogin) (*response.User, error)
	IsEmailAvailable(ctx context.Context, email string) (bool, error)
	UploadAvatar(ctx context.Context, userId int, file *multipart.FileHeader) (*response.User, error)
}
