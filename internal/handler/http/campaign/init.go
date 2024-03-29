package campaign

import (
	"bwa-startup/internal/common"
	"bwa-startup/internal/common/errors"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/service/campaign"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type handlerImpl struct {
	service campaign.Service
}

func (h *handlerImpl) GetAllCampaign(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(response.New{
		Success: true,
		Code:    http.StatusOK,
		Message: "belum di develop",
	})
}

func (h *handlerImpl) GetCampaign(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))
	if userId == -1 {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success: false, Code: http.StatusBadRequest, Message: "Invalid Request",
		})

	}

	var resp response.New
	campaignByUserId, err := h.service.GetCampaignByUserId(c.Context(), userId)
	if err != nil {
		resp = response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	} else {
		resp = response.New{
			Success: false,
			Code:    http.StatusOK,
			Message: "success",
			Data:    campaignByUserId,
		}
	}

	return c.Status(resp.Code).JSON(resp)
}

func (h *handlerImpl) GetCampaignById(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))
	if userId == -1 {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success: false, Code: http.StatusBadRequest, Message: "Invalid Request",
		})

	}

	campaignIdString := c.Params("campaignId")
	campaignId, err := strconv.Atoi(campaignIdString)

	campaignDetail, err := h.service.GetCampaignDetailById(c.Context(), userId, campaignId)

	var resp response.New
	if err != nil {
		resp = response.New{
			Success: false, Code: errors.StatusCode(err.Error()), Message: err.Error(),
		}
	} else {
		resp = response.New{
			Success: true, Code: http.StatusOK, Message: "success", Data: campaignDetail,
		}
	}

	return c.Status(resp.Code).JSON(resp)
}

func (h *handlerImpl) CreateCampaign(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))
	if userId == -1 {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "Invalid Request"})

	}

	//request
	body := request.Campaign{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{
			Success: false, Code: http.StatusBadRequest, Message: err.Error(),
		})

	}
	body.UserId = userId

	createCampaign, err := h.service.CreateCampaign(c.Context(), &body)
	var resp response.New
	if err != nil {
		resp = response.New{
			Success: false, Code: http.StatusInternalServerError, Message: err.Error(),
		}
	} else {
		resp = response.New{
			Success: true, Code: http.StatusOK, Message: "success", Data: createCampaign,
		}
	}

	return c.Status(resp.Code).JSON(resp)
}

func (h *handlerImpl) UpdateCampaign(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))
	if userId == -1 {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "Invalid Request"})
	}

	campaignId, err := strconv.Atoi(c.Params("campaignId"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{
			Success: false, Code: http.StatusBadRequest, Message: err.Error(),
		})

	}

	//request
	body := request.Campaign{}
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{
			Success: false, Code: http.StatusBadRequest, Message: err.Error(),
		})
	}
	body.UserId = userId

	updateCampaign, err := h.service.UpdateCampaign(c.Context(), campaignId, &body)
	var resp response.New
	if err != nil {
		resp = response.New{
			Success: false, Code: http.StatusInternalServerError, Message: err.Error(),
		}
	} else {
		resp = response.New{
			Success: true, Code: http.StatusOK, Message: "success", Data: updateCampaign,
		}
	}

	return c.Status(resp.Code).JSON(resp)

}

func (h *handlerImpl) UploadImage(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))
	if userId == -1 {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid data"})
	}

	campaignId, err := strconv.Atoi(c.Params("campaignId"))

	isPrimary, err := strconv.ParseBool(c.FormValue("is_primary", "true"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "Invalid Request", ErrorDetail: err})

	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "Invalid Request", ErrorDetail: err})
	}

	err = h.service.UploadImage(c.Context(), userId, campaignId, file, isPrimary)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "Invalid Request", ErrorDetail: err})

	}

	return c.Status(http.StatusOK).JSON(response.New{
		Success: true,
		Code:    http.StatusOK,
		Message: strconv.Itoa(userId),
		Data:    campaignId,
	})
}

func NewHandler(service campaign.Service) Handler {
	return &handlerImpl{service: service}
}
