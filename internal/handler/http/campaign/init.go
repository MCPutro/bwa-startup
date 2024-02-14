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

	campaignByUserId, err := h.service.GetCampaignByUserId(c.Request.Context(), unitID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success:      false,
			Code:         http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.New{
		Success: true,
		Code:    http.StatusOK,
		Data:    campaignByUserId,
	})
}

func NewHandler(service campaign.Service) Handler {
	return &handlerImpl{service: service}
}
