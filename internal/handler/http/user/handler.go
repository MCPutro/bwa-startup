package user

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
	CheckEmailAvailable(c *gin.Context)
	UploadAvatar(c *gin.Context)
}
