package user

import (
	"bwa-startup/config"
	"bwa-startup/internal/common"
	newError "bwa-startup/internal/common/errors"
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
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	user     user.Repository
	config   config.Config
	auth     auth.Repository
	firebase firebase.Repository
}

func (us *userServiceImpl) GetAll(ctx context.Context) ([]*entity.User, error) {
	// resp = new([]entity.User)

	u, err := us.user.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *userServiceImpl) Register(ctx context.Context, newUser *request.RegisterUser) (*response.User, error) {
	//check email is already register
	isExistingUser, err := us.IsEmailAvailable(ctx, newUser.Email)
	if err != nil {
		return nil, err
	}
	if !isExistingUser {
		return nil, newError.ErrEmailAlreadyRegister
	}

	//save new user to database
	userEntity := newUser.ToEntity()
	u, err := us.user.Create(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	//generate token
	token, err := us.auth.GenerateToken(u)
	if err != nil {
		return nil, err
	}

	return u.ToUserResponse( //us.config.FirebaseConf().BucketName(),
		token), nil
}

func (us *userServiceImpl) Login(ctx context.Context, newUser *request.UserLogin) (*response.User, error) {
	existingUser, err := us.user.FindByEmail(ctx, newUser.Email)
	if err != nil {
		return nil, err
	}

	// if user is found, then check password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(newUser.Password))
	if err != nil {
		return nil, newError.ErrEmailAndPasswordNotMatch
	}

	//generate token
	token, err := us.auth.GenerateToken(existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser.ToUserResponse(token), nil
}

func (us *userServiceImpl) IsEmailAvailable(ctx context.Context, email string) (bool, error) {
	//get user by email
	existingUser, err := us.user.FindByEmail(ctx, email)
	if errors.Is(err, newError.ErrEmailNotFound) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	if existingUser != nil && existingUser.ID > 0 {
		return false, nil
	}

	return true, nil
}

func (us *userServiceImpl) UploadAvatar(ctx context.Context, userId int, file *multipart.FileHeader) (*response.User, error) {
	//validate file type from content type

	if err := common.IsSupportedImageType(us.config.ImageConf().SupportType(), file.Header.Get("Content-Type")); err != nil {
		return nil, err
	}

	//check user id
	existingUser, err := us.user.FindById(ctx, userId)
	if err != nil || existingUser == nil {
		return nil, err
	}

	//open file
	bufferFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer bufferFile.Close()

	//upload file
	imagePath := fmt.Sprint(us.firebase.BucketPath(), "/users/", userId, "/avatar/", userId, "-avatar", filepath.Ext(file.Filename))
	token, err := us.firebase.UploadFile(ctx, bufferFile, imagePath)
	if err != nil {
		return nil, err
	}

	//update user
	existingUser.Image = common.GetUrlImage(us.firebase.BucketName(), imagePath, token) //imagePath
	updatedUser, err := us.user.Update(ctx, existingUser)
	if err != nil {
		return nil, err
	}

	return updatedUser.ToUserResponse(""), nil
}

func NewService(conf config.Config, user user.Repository, auth auth.Repository, firebase firebase.Repository) Service {
	return &userServiceImpl{
		user:     user,
		config:   conf,
		auth:     auth,
		firebase: firebase,
	}
}
