package routes

import (
	"bwa-startup/internal/handler/http/campaign"
	"github.com/gin-gonic/gin"
)

func CampaignRoute(group *gin.RouterGroup, handler campaign.Handler, middleware gin.HandlerFunc) {
	route := group.Group("/campaign")

	route.GET("", middleware, handler.GetCampaign)
	route.GET("/:campaignId", middleware, handler.GetCampaignById)
}
