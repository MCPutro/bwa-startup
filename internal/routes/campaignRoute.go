package routes

import (
	"bwa-startup/internal/handler/http/campaign"

	"github.com/gofiber/fiber/v2"
)

func CampaignRoute(group fiber.Router, handler campaign.Handler, middleware fiber.Handler) {
	route := group.Group("/campaign")

	route.Use(middleware)

	route.Get("/", handler.GetCampaign)
	route.Get("/:campaignId", handler.GetCampaignById)
	route.Post("/", handler.CreateCampaign)
	route.Put("/:campaignId", handler.UpdateCampaign)
	route.Post("/:campaignId/image", handler.UploadImage)
	route.Get("/:campaignId/transactions", handler.FindTrx)
}
