package entity

import (
	"time"
)

type Transaction struct {
	ID         int `gorm:"primarykey"`
	CampaignId int
	UserId     int
	User       User
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}
