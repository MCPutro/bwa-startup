package campaign

import "github.com/gin-gonic/gin"

type Handler interface {
	GetAllCampaign(c *gin.Context)
	GetCampaign(c *gin.Context)
	GetCampaignById(c *gin.Context)
	CreateCampaign(c *gin.Context)
}
