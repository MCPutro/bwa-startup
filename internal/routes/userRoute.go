package routes

import (
	"bwa-startup/internal/handler/http/user"
	"github.com/gin-gonic/gin"
)

func UserRoutePublic(group *gin.RouterGroup, handler user.Handler) {
	route := group.Group("/user")

	route.POST("/register", handler.RegisterUser)
	route.POST("/login", handler.Login)
	route.POST("/emailCheckers", handler.CheckEmailAvailable)

}

func UserRoutePrivate(group *gin.RouterGroup, handler user.Handler, middleware gin.HandlerFunc) {
	route := group.Group("/user")

	route.POST("/uploadAvatar", middleware, handler.UploadAvatar)
}
