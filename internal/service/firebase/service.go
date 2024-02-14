package firebase

import (
	"bwa-startup/internal/handler/response"
	"context"
	"mime/multipart"
)

type Service interface {
	UploadImage(ctx context.Context, userId int, uploadedFile multipart.File, uploadedFileHeader *multipart.FileHeader) (*response.User, error)
}
