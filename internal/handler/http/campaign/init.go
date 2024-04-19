package campaign

import (
	"bwa-startup/internal/common"
	"bwa-startup/internal/common/errors"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/service/campaign"
	"bwa-startup/internal/service/transaction"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type handlerImpl struct {
	campaign campaign.Service
	trx      transaction.Service
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
			Success: false, Code: http.StatusBadRequest, Message: "invalid request",
		})

	}

	var resp response.New
	campaignByUserId, err := h.campaign.GetByUserId(c.Context(), userId)
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
			Success: false, Code: http.StatusBadRequest, Message: "invalid request",
		})

	}

	campaignIdString := c.Params("campaignId")
	campaignId, err := strconv.Atoi(campaignIdString)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid request", ErrorDetail: err})
	}

	campaignDetail, err := h.campaign.GetDetailById(c.Context(), userId, campaignId)

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
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid request"})

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

	createCampaign, err := h.campaign.Save(c.Context(), &body)
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
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid request"})
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

	updateCampaign, err := h.campaign.Update(c.Context(), campaignId, &body)
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
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid request", ErrorDetail: err.Error()})
	}

	isPrimary, err := strconv.ParseBool(c.FormValue("is_primary", "true"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid request", ErrorDetail: err.Error()})

	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid request", ErrorDetail: err.Error()})
	}

	err = h.campaign.UploadImage(c.Context(), userId, campaignId, file, isPrimary)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid request", ErrorDetail: err.Error()})

	}

	return c.Status(http.StatusOK).JSON(response.New{
		Success: true,
		Code:    http.StatusOK,
		Message: "success",
		Data:    campaignId,
	})
}

func (h *handlerImpl) FindTrx(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))
	if userId == -1 {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid data"})
	}

	campaignId, err := strconv.Atoi(c.Params("campaignId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: "invalid request", ErrorDetail: err.Error()})
	}

	result, err := h.trx.FindByCampaignId(c.Context(), userId, campaignId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{Success: false, Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(response.New{
		Success: true,
		Code:    http.StatusOK,
		Message: "success",
		Data:    result,
	})
}

func NewHandler(campaign campaign.Service, trx transaction.Service) Handler {
	return &handlerImpl{
		campaign: campaign,
		trx:      trx,
	}
}
