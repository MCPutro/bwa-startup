package firebase

import (
	"context"
	"mime/multipart"
)

type Repository interface {
	//UploadImage(ctx context.Context, userId int, uploadedFile multipart.File, uploadedFileHeader *multipart.FileHeader) (*response.User, error)
	UploadFile(ctx context.Context, file multipart.File, patch string) (string, error)
	BucketName() string
	BucketPath() string
}
