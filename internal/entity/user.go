package entity

import (
	"bwa-startup/internal/handler/response"
	"fmt"
	"net/url"
	"time"
)

type User struct {
	ID         int
	Name       string
	Occupation string
	Email      string
	Password   string
	Image      string
	ImageToken string
	Role       string
	CreatedAt  time.Time `sql:"type:timestamp"`
	UpdatedAt  time.Time `sql:"type:timestamp"`
}

func (u *User) ToUserResponse(bucket string, token string) *response.User {
	return &response.User{
		ID:         u.ID,
		Name:       u.Name,
		Occupation: u.Occupation,
		Avatar:     u.GetUrlAvatar(bucket),
		Email:      u.Email,
		Token:      token,
	}
}

func (u *User) GetUrlAvatar(bucket string) string {
	var urlAvatar string
	if u.Image != "" && u.ImageToken != "" {
		urlAvatar = fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", bucket, url.PathEscape(u.Image), u.ImageToken)
	}
	return urlAvatar
}

func (u *User) TableName() string {
	return "users"
}
