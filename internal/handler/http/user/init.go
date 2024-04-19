package user

import (
	"bwa-startup/config"
	"bwa-startup/internal/common"
	newError "bwa-startup/internal/common/errors"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/service/transaction"
	"bwa-startup/internal/service/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handlerImpl struct {
	user  user.Service
	image config.ImageConf
	trx   transaction.Service
}

func (h *handlerImpl) Login(c *fiber.Ctx) error {
	var body request.UserLogin
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(newError.StatusCode(err.Error())).JSON(response.New{
			Success: false,
			Code:    newError.StatusCode(err.Error()),
			Message: err.Error(),
		})

	}

	u, err := h.user.Login(c.Context(), &body)
	if err != nil {
		return c.Status(newError.StatusCode(err.Error())).JSON(response.New{
			Success: false,
			Code:    newError.StatusCode(err.Error()),
			Message: err.Error(),
		})

	}

	return c.Status(http.StatusOK).JSON(response.New{
		Success: true,
		Code:    http.StatusOK,
		Data:    u,
	})
}

func (h *handlerImpl) RegisterUser(c *fiber.Ctx) error {
	body := request.RegisterUser{}

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	u, err := h.user.Register(c.Context(), &body)
	if err != nil {
		return c.Status(newError.StatusCode(err.Error())).JSON(response.New{
			Success: false,
			Code:    newError.StatusCode(err.Error()),
			Message: err.Error(),
		})

	}

	return c.Status(http.StatusCreated).JSON(response.New{
		Success: true,
		Code:    http.StatusCreated,
		Data:    u,
	})

}

func (h *handlerImpl) CheckEmailAvailable(c *fiber.Ctx) error {
	var body request.UserLogin

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	b, err := h.user.IsEmailAvailable(c.Context(), body.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	if b {
		return c.Status(http.StatusOK).JSON(response.New{
			Success: true,
			Code:    http.StatusOK,
			Message: "email available",
		})
	} else {
		return c.Status(http.StatusNotAcceptable).JSON(response.New{
			Success: false,
			Code:    http.StatusNotAcceptable,
			Message: "email already use",
		})
	}
}

func (h *handlerImpl) UploadAvatar(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))

	//get file from req

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success:     false,
			Code:        http.StatusBadRequest,
			Message:     err.Error(),
			ErrorDetail: nil,
			Data:        nil,
		})

	}

	//check file size
	if file.Size >= h.image.MaxAvatarSize() {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success:     false,
			Code:        http.StatusBadRequest,
			Message:     "file size limit exceeded",
			ErrorDetail: nil,
			Data:        nil,
		})
	}

	//upload file
	resp, err := h.user.UploadAvatar(c.Context(), userId, file)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})

	}

	return c.Status(http.StatusOK).JSON(response.New{
		Success: true,
		Code:    http.StatusOK,
		Data:    resp,
	})
}

func (h *handlerImpl) FindTrx(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))
	if userId <= 0 {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
		})
	}

	//get trx list
	trx, err := h.trx.FindByUserId(c.Context(), userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(response.New{
		Success: true,
		Code:    http.StatusOK,
		Data:    trx,
	})
}

func (h *handlerImpl) CreateTrx(c *fiber.Ctx) error {
	userId := common.GetUserId(c.Locals("userID"))
	if userId <= 0 {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "invalid user id",
		})
	}

	var body request.Transaction

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	body.UserId = userId

	newTrx, err := h.trx.Save(c.Context(), &body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(response.New{
		Success: true,
		Code:    http.StatusCreated,
		Data:    newTrx,
	})
}

func NewHandler(service user.Service, image config.ImageConf, trx transaction.Service) Handler {
	return &handlerImpl{
		user:  service,
		image: image,
		trx:   trx,
	}
}
