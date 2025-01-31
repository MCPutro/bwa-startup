package entity

import (
	"bwa-startup/internal/handler/response"
	"time"
)

type User struct {
	ID         int       `gorm:"primary_key;column:id;<-:create"`
	Name       string    `gorm:"column:name"`
	Occupation string    `gorm:"column:occupation"`
	Email      string    `gorm:"column:email"`
	Password   string    `gorm:"column:password"`
	Image      string    `gorm:"column:image"`
	Role       string    `gorm:"column:role"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) ToUserResponse(token string) *response.User {
	return &response.User{
		ID:         u.ID,
		Name:       u.Name,
		Occupation: u.Occupation,
		Avatar:     u.Image, //u.GetUrlAvatar(bucket),
		Email:      u.Email,
		Token:      token,
	}
}

func (u *User) TableName() string {
	return "users"
}
