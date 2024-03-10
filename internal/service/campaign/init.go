package campaign

import (
	"bwa-startup/config"
	"bwa-startup/internal/common"
	"bwa-startup/internal/entity"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/campaign"
	"bwa-startup/internal/repository/firebase"
	"bwa-startup/internal/repository/user"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"
)

type campaignServiceImpl struct {
	campaign     campaign.Repository
	user         user.Repository
	firebase     firebase.Repository
	config       config.Config
	firebaseConf config.FirebaseConfig
}

func (s *campaignServiceImpl) GetCampaignByUserId(ctx context.Context, userId int) ([]*response.Campaign, error) {
	//check user id
	_, err := s.user.FindById(ctx, userId)
	if err == nil {
		return nil, errors.New("user id not found")
	}

	campaignByUserId, err := s.campaign.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return campaignByUserId.ToCampaignRespList(), nil
}

func (s *campaignServiceImpl) GetCampaignDetailById(ctx context.Context, userId, campaignId int) (*response.CampaignDetail, error) {
	if userId <= 0 {
		return nil, nil
	}

	campaignById, err := s.campaign.FindById(ctx, userId, campaignId)
	if err != nil {
		return nil, err
	}

	return campaignById.ToCampaignDetailResp(s.firebaseConf.BucketName()), nil
}

func (s *campaignServiceImpl) CreateCampaign(ctx context.Context, campaign *request.Campaign) (*response.CampaignDetail, error) {
	//check user
	existingUser, _ := s.user.FindById(ctx, campaign.UserId)
	if existingUser == nil {
		return nil, errors.New("user id not found")
	}

	save, err := s.campaign.Save(ctx, campaign.ToEntity())
	if err != nil {
		return nil, err
	}

	save.User = *existingUser

	return save.ToCampaignDetailResp(s.firebaseConf.BucketName()), nil
}

func (s *campaignServiceImpl) UpdateCampaign(ctx context.Context, campaignId int, newCampaign *request.Campaign) (*response.CampaignDetail, error) {
	//get existing campaign
	existingCampaign, _ := s.campaign.FindById(ctx, newCampaign.UserId, campaignId)
	if existingCampaign == nil {
		return nil, errors.New("campaign not found")
	}

	existingCampaign.Name = newCampaign.Name
	existingCampaign.SortDescription = newCampaign.ShortDescription
	existingCampaign.Description = newCampaign.Description
	existingCampaign.Perks = newCampaign.Perks
	existingCampaign.GoalAmount = newCampaign.GoalAmount
	existingCampaign.UpdatedAt = time.Now()

	updatedCampaign, err := s.campaign.Update(ctx, existingCampaign)
	if err != nil {
		return nil, err
	}

	return updatedCampaign.ToCampaignDetailResp(s.firebaseConf.BucketName()), nil
}

func (s *campaignServiceImpl) UploadImage(ctx context.Context, userId, campaignId int, file *multipart.FileHeader, isPrimary bool) error {
	//check content type
	if err := common.IsSupportedImageType(s.config.ImageConf().SupportType(), file.Header.Get("Content-Type")); err != nil {
		return err
	}

	//check user id and campaign id
	existingCampaign, _ := s.campaign.FindById(ctx, userId, campaignId)
	if existingCampaign == nil {
		return errors.New("campaign not found")
	}

	//upload to firebase
	bufferFile, err := file.Open()
	if err != nil {
		return err
	}
	defer bufferFile.Close()

	imagePath := fmt.Sprint(s.firebaseConf.BucketPath(), "/campaigns/", campaignId, "/", file.Filename)
	tokenFile, err := s.firebase.UploadFile(ctx, bufferFile, imagePath)
	if err != nil {
		return err
	}

	//update database
	campaignImage := entity.CampaignImage{
		CampaignID: campaignId,
		Filename:   imagePath + "?alt=media&token=" + tokenFile,
		IsPrimary:  isPrimary,
	}
	err = s.campaign.CreateImage(ctx, &campaignImage)
	if err != nil {
		return err
	}

	return nil
}

func NewService(cfg config.Config, repo campaign.Repository, user user.Repository, firebase firebase.Repository) Service {
	return &campaignServiceImpl{
		campaign:     repo,
		user:         user,
		firebase:     firebase,
		config:       cfg,
		firebaseConf: cfg.FirebaseConf(),
	}
}
