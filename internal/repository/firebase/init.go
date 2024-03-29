package firebase

import (
	"bwa-startup/config"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	cloudFirebase "firebase.google.com/go/v4"
	firebaseStorage "firebase.google.com/go/v4/storage"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

type firebaseRepositoryImpl struct {
	cfg     config.FirebaseConfig
	app     *cloudFirebase.App
	storage *firebaseStorage.Client
}

func (fs *firebaseRepositoryImpl) UploadFile(ctx context.Context, file multipart.File, patch string) (string, error) {
	bucket, err := fs.storage.Bucket(fs.cfg.BucketName())
	if err != nil {
		return "", err
	}

	// upload it to google cloud store
	wc := bucket.Object(patch).NewWriter(ctx)
	//Set the attribute
	id := uuid.New()
	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}
	err = wc.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}

	return id.String(), nil
}

//// UploadImage implements Service.
//func (fs *firebaseRepositoryImpl) UploadImage(ctx context.Context, userId int, file multipart.File, fileHeader *multipart.FileHeader) (*response.User, error) {
//	//validate file type
//	contentType := fileHeader.Header.Get("Content-Type")
//	splitContentType := strings.Split(contentType, "/")
//
//	if strings.ToUpper(splitContentType[0]) != "IMAGE" && !fs.supportImageType[strings.ToLower(splitContentType[1])] {
//		return nil, errors.New("unsupported image type")
//	}
//
//	//check user id
//	existingUser, err := fs.userRepo.FindById(ctx, userId)
//	if err != nil || existingUser == nil {
//		return nil, err
//	}
//
//	bucket, err := fs.storage.Bucket(fs.cfg.BucketName())
//	if err != nil {
//		return nil, err
//	}
//
//	imagePath := fmt.Sprint(fs.cfg.BucketPath(), "/users/", userId, "/avatar/", userId, "-avatar.", strings.Split(fileHeader.Filename, ".")[1])
//
//	// upload it to google cloud store
//	wc := bucket.Object(imagePath).NewWriter(ctx)
//	//Set the attribute
//	id := uuid.New()
//	wc.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}
//	if _, err = io.Copy(wc, file); err != nil {
//		return nil, err
//	}
//	err = wc.Close()
//	if err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	//update database
//	existingUser.Image = imagePath
//	existingUser.ImageToken = id.String()
//	updatedUser, err := fs.userRepo.Update(ctx, existingUser)
//	if err != nil {
//		return nil, err
//	}
//
//	//return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", fs.cfg.BucketName(), url.PathEscape(imagePath), id.String()), nil
//	return updatedUser.ToUserResponse(fs.cfg.BucketName(), ""), nil
//}

func resizeImage() {
	//image.Rect()
}

// BucketName implements Repository.
func (fs *firebaseRepositoryImpl) BucketName() string {
	return fs.cfg.BucketName()
}

// BucketPath implements Repository.
func (fs *firebaseRepositoryImpl) BucketPath() string {
	return fs.cfg.BucketPath()
}

func NewRepository(cfg config.Config) Repository {

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

	return &firebaseRepositoryImpl{
		cfg:     firebaseConf,
		app:     nApp,
		storage: storages,
	}
}
