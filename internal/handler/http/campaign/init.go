package campaign

import (
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/service/campaign"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handlerImpl struct {
	service campaign.Service
}

func (h *handlerImpl) GetAllCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, response.New{
		Success: true,
		Code:    http.StatusOK,
		Message: "belum di develop",
	})
}

func (h *handlerImpl) GetCampaign(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := fmt.Sprint(userIDInterface)
	unitID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	campaignByUserId, err := h.service.GetCampaignByUserId(c.Request.Context(), unitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.New{
		Success: true,
		Code:    http.StatusOK,
		Data:    campaignByUserId,
	})
}

func (h *handlerImpl) GetCampaignById(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userIdString := fmt.Sprint(userIDInterface)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	campaignIdString := c.Param("campaignId")
	campaignId, err := strconv.Atoi(campaignIdString)

	campaignDetail, err := h.service.GetCampaignDetailById(c.Request.Context(), userId, campaignId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	c.JSON(http.StatusOK, response.New{
		Success: true,
		Code:    http.StatusOK,
		Data:    campaignDetail,
	})
}

func NewHandler(service campaign.Service) Handler {
	return &handlerImpl{service: service}
}
