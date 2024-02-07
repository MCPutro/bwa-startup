package firebase

import (
	"bwa-startup/config"
	"bwa-startup/internal/repository/user"
	"context"
	"encoding/base64"
	cloudFirebase "firebase.google.com/go/v4"
	firebaseStorage "firebase.google.com/go/v4/storage"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"log"
	"mime/multipart"
	"net/url"
)

type Service interface {
	UploadImage(ctx context.Context, userId int, uploadedFile multipart.File, uploadedFileHeader *multipart.FileHeader) (string, error)
}

type serviceImpl struct {
	cfg      config.FirebaseConfig
	app      *cloudFirebase.App
	userRepo user.Repository
	storage  *firebaseStorage.Client
}

// UploadImage implements Service.
func (fs *serviceImpl) UploadImage(ctx context.Context, userId int, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	//check user id
	existingUser, err := fs.userRepo.FindById(ctx, userId)
	if err != nil {
		return "", err
	}

	bucket, err := fs.storage.Bucket(fs.cfg.BucketName())
	if err != nil {
		return "", err
	}

	imagePath := fmt.Sprint(fs.cfg.BucketPath(), "/users/", userId, "/avatar/", userId, "-", fileHeader.Filename)

	// upload it to google cloud store
	wc := bucket.Object(imagePath).NewWriter(ctx)
	//Set the attribute
	id := uuid.New()
	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	if _, err = io.Copy(wc, file); err != nil {
		log.Println(err)
		return "", err
	}
	err = wc.Close()
	if err != nil {
		log.Println(err)
	}

	//update database
	existingUser.Image = imagePath
	existingUser.ImageToken = id.String()
	_, err = fs.userRepo.Update(ctx, existingUser)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", fs.cfg.BucketName(), url.PathEscape(imagePath), id.String()), nil
}

func resizeImage() {
	//image.Rect()
}

func NewService(cfg config.FirebaseConfig, repo user.Repository) Service {

	ctx := context.Background()

	key, err := base64.StdEncoding.DecodeString(cfg.AdminKey())
	if err != nil {
		return nil
	}

	opt := option.WithCredentialsJSON(key)
	nApp, err := cloudFirebase.NewApp(ctx, &cloudFirebase.Config{
		StorageBucket:    fmt.Sprint("gs://", cfg.BucketName()),
		ProjectID:        cfg.ProjectId(),
		ServiceAccountID: cfg.ServiceAccount(),
	}, opt)

	if err != nil {
		log.Println("err : ", err)
		return nil
	}

	storages, err := nApp.Storage(ctx)
	if err != nil {
		log.Println("err : ", err)
		return nil
	}

	return &serviceImpl{
		cfg:      cfg,
		app:      nApp,
		userRepo: repo,
		storage:  storages,
	}
}
