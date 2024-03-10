package campaign

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetAllCampaign(c *fiber.Ctx) error
	GetCampaign(c *fiber.Ctx) error
	GetCampaignById(c *fiber.Ctx) error
	CreateCampaign(c *fiber.Ctx) error
	UpdateCampaign(c *fiber.Ctx) error
	UploadImage(c *fiber.Ctx) error
}
