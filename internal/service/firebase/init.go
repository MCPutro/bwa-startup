package firebase

import (
	"bwa-startup/config"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/user"
	"context"
	"encoding/base64"
	"errors"
	cloudFirebase "firebase.google.com/go/v4"
	firebaseStorage "firebase.google.com/go/v4/storage"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"io"
	"log"
	"mime/multipart"
	"strings"
)

type serviceImpl struct {
	cfg              config.FirebaseConfig
	userRepo         user.Repository
	supportImageType map[string]bool
	app              *cloudFirebase.App
	storage          *firebaseStorage.Client
}

// UploadImage implements Service.
func (fs *serviceImpl) UploadImage(ctx context.Context, userId int, file multipart.File, fileHeader *multipart.FileHeader) (*response.User, error) {
	//validate file type
	contentType := fileHeader.Header.Get("Content-Type")
	splitContentType := strings.Split(contentType, "/")

	if strings.ToUpper(splitContentType[0]) != "IMAGE" && !fs.supportImageType[strings.ToLower(splitContentType[1])] {
		return nil, errors.New("unsupported image type")
	}

	//check user id
	existingUser, err := fs.userRepo.FindById(ctx, userId)
	if err != nil || existingUser == nil {
		return nil, err
	}

	bucket, err := fs.storage.Bucket(fs.cfg.BucketName())
	if err != nil {
		return nil, err
	}

	imagePath := fmt.Sprint(fs.cfg.BucketPath(), "/users/", userId, "/avatar/", userId, "-avatar.", strings.Split(fileHeader.Filename, ".")[1])

	// upload it to google cloud store
	wc := bucket.Object(imagePath).NewWriter(ctx)
	//Set the attribute
	id := uuid.New()
	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	if _, err = io.Copy(wc, file); err != nil {
		return nil, err
	}
	err = wc.Close()
	if err != nil {
		log.Println(err)
	}

	//update database
	existingUser.Image = imagePath
	existingUser.ImageToken = id.String()
	updatedUser, err := fs.userRepo.Update(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	//return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", fs.cfg.BucketName(), url.PathEscape(imagePath), id.String()), nil
	userResponse := updatedUser.ToUserResponse(fs.cfg.BucketName(), "")

	return &userResponse, nil
}

func resizeImage() {
	//image.Rect()
}

func NewService(cfg config.Config, repo user.Repository) Service {

	ctx := context.Background()

	firebaseConf := cfg.FirebaseConf()

	key, err := base64.StdEncoding.DecodeString(firebaseConf.AdminKey())
	if err != nil {
		return nil
	}

	opt := option.WithCredentialsJSON(key)
	nApp, err := cloudFirebase.NewApp(ctx, &cloudFirebase.Config{
		StorageBucket:    fmt.Sprint("gs://", firebaseConf.BucketName()),
		ProjectID:        firebaseConf.ProjectId(),
		ServiceAccountID: firebaseConf.ServiceAccount(),
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
		cfg:              firebaseConf,
		supportImageType: cfg.ImageSupport(),
		app:              nApp,
		userRepo:         repo,
		storage:          storages,
	}
}
