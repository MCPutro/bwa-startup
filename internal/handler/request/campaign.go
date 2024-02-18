package request

import (
	"bwa-startup/internal/entity"
)

type Campaign struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	GoalAmount       int    `json:"goal_amount"`
	Perks            string `json:"perks"`
	UserId           int    `json:"-"`
}

func (c *Campaign) ToEntity() *entity.Campaign {
	return &entity.Campaign{
		UserId:          c.UserId,
		Name:            c.Name,
		SortDescription: c.ShortDescription,
		Description:     c.Description,
		Perks:           c.Perks,
		GoalAmount:      c.GoalAmount,
	}
}
