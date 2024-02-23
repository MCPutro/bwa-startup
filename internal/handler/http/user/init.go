package user

import (
	"bwa-startup/config"
	"bwa-startup/internal/handler/request"
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/service/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerImpl struct {
	service user.Service
	image   config.ImageConf
}

// Login implements Handler.
func (h *handlerImpl) Login(c *gin.Context) {
	var body request.UserLogin
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	u, err := h.service.Login(c.Request.Context(), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.New{
		Success: true,
		Code:    http.StatusOK,
		Data:    u,
	})
}

// RegisterUser implements Handler.
func (h *handlerImpl) RegisterUser(c *gin.Context) {
	body := request.RegisterUser{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	u, err := h.service.Register(c.Request.Context(), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.New{
		Success: true,
		Code:    http.StatusCreated,
		Data:    u,
	})

}

// CheckEmailAvailable implements Handler.
func (h *handlerImpl) CheckEmailAvailable(c *gin.Context) {
	var body request.UserLogin

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	b, err := h.service.IsEmailAvailable(c.Request.Context(), body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if b {
		c.JSON(http.StatusCreated, response.New{
			Success: true,
			Code:    http.StatusOK,
			Message: "email available",
		})
	} else {
		c.JSON(http.StatusCreated, response.New{
			Success: false,
			Code:    http.StatusNotAcceptable,
			Message: "email already use",
		})
	}
}

// UploadAvatar implements Handler.
func (h *handlerImpl) UploadAvatar(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := fmt.Sprint(userIDInterface)
	unitID, err := strconv.Atoi(userID)

	//get file from req
	file, uploadedFileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	defer file.Close()

	//check file size
	if uploadedFileHeader.Size >= h.image.MaxAvatarSize() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "file size limit exceeded",
			"error":   true,
		})
		return
	}

	resp, err := h.service.UploadAvatar(c.Request.Context(), unitID, file, uploadedFileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.New{
			Success: false,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.New{
		Success: true,
		Code:    http.StatusOK,
		Data:    resp,
	})
}

func NewHandler(service user.Service, image config.ImageConf) Handler {
	return &handlerImpl{
		service: service,
		image:   image,
	}
}
