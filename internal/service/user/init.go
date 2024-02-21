package user

import (
	"bwa-startup/config"
	"bwa-startup/internal/entity"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/auth"
	"bwa-startup/internal/repository/firebase"
	"bwa-startup/internal/repository/user"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	user     user.Repository
	config   config.Config
	auth     auth.Repository
	firebase firebase.Repository
}

// GetAll implements Service.
func (us *userServiceImpl) GetAll(ctx context.Context) ([]*entity.User, error) {
	// resp = new([]entity.User)

	u, err := us.user.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Register implements UserService.
func (us *userServiceImpl) Register(ctx context.Context, newUser *request.RegisterUser) (*response.User, error) {
	//check email is already register
	isExistingUser, err := us.IsEmailAvailable(ctx, newUser.Email)
	if err != nil {
		return nil, err
	}
	if !isExistingUser {
		return nil, errors.New("email already register")
	}

	//save new user to database
	userEntity := newUser.ToEntity()
	u, err := us.user.Save(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	//generate token
	token, err := us.auth.GenerateToken(u)
	if err != nil {
		return nil, err
	}

	return u.ToUserResponse(us.config.FirebaseConf().BucketName(), token), nil
}

// Login implements Service.
func (us *userServiceImpl) Login(ctx context.Context, newUser *request.UserLogin) (*response.User, error) {
	existingUser, err := us.user.FindByEmail(ctx, newUser.Email)
	if err != nil {
		return nil, err
	}

	// if user is found, then check password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(newUser.Password))
	if err != nil {
		return nil, errors.New("email and password not match")
	}

	//generate token
	token, err := us.auth.GenerateToken(existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser.ToUserResponse(us.config.FirebaseConf().BucketName(), token), nil
}

// IsEmailAvailable implements Service.
func (us *userServiceImpl) IsEmailAvailable(ctx context.Context, email string) (bool, error) {
	//get user by email
	existingUser, err := us.user.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	if existingUser != nil && existingUser.ID > 0 {
		return false, nil
	}

	return true, nil
}

// UploadAvatar implements Service.
func (us *userServiceImpl) UploadAvatar(ctx context.Context, userId int, uploadedFile multipart.File, uploadedFileHeader *multipart.FileHeader) (*response.User, error) {
	//validate file type
	contentType := uploadedFileHeader.Header.Get("Content-Type")
	splitContentType := strings.Split(contentType, "/")
	if strings.ToUpper(splitContentType[0]) != "IMAGE" && !us.config.ImageConf().SupportType(strings.ToLower(splitContentType[1])) {
		return nil, errors.New("unsupported image type")
	}

	//check user id
	existingUser, err := us.user.FindById(ctx, userId)
	if err != nil || existingUser == nil {
		return nil, err
	}

	//upload file
	imagePath := fmt.Sprint(us.config.FirebaseConf().BucketPath(), "/users/", userId, "/avatar/", userId, "-avatar.", strings.Split(uploadedFileHeader.Filename, ".")[1])
	token, err := us.firebase.UploadFile(ctx, uploadedFile, imagePath)
	if err != nil {
		return nil, err
	}

	//update user
	existingUser.Image = imagePath
	existingUser.ImageToken = token
	updatedUser, err := us.user.Update(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	return updatedUser.ToUserResponse(us.config.FirebaseConf().BucketName(), ""), nil
}

func NewService(conf config.Config, user user.Repository, auth auth.Repository, firebase firebase.Repository) Service {
	return &userServiceImpl{
		user:     user,
		config:   conf,
		auth:     auth,
		firebase: firebase,
	}
}
