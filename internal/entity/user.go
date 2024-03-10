package entity

import (
	"bwa-startup/internal/handler/response"
	"fmt"
	"net/url"
	"time"
)

type User struct {
	ID         int       `gorm:"primary_key;column:id;<-:create"`
	Name       string    `gorm:"column:name"`
	Occupation string    `gorm:"column:occupation"`
	Email      string    `gorm:"column:email"`
	Password   string    `gorm:"column:password"`
	Image      string    `gorm:"column:image"`
	ImageToken string    `gorm:"column:image_token"`
	Role       string    `gorm:"column:role"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
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
