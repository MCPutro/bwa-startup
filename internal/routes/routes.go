package routes

import (
	"bwa-startup/internal/handler/http/campaign"
	"bwa-startup/internal/handler/http/user"
	"bwa-startup/internal/middleware"
	"bwa-startup/internal/repository"
	"bwa-startup/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine, service service.Service, repo repository.Repository) {
	route := r.Group("/")
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	apiPublic := r.Group("/api/v1")

	userHandler := user.NewHandler(service.UserService())

	authMiddleware := middleware.New(repo.AuthRepository())

	UserRoutePublic(apiPublic, userHandler)

	//with middleware
	UserRoutePrivate(apiPublic, userHandler, authMiddleware)

	campaignHandler := campaign.NewHandler(service.CampaignService())
	CampaignRoute(apiPublic, campaignHandler, authMiddleware)
}
