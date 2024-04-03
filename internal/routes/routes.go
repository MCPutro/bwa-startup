package routes

import (
	"bwa-startup/config"
	"bwa-startup/internal/app"
	"bwa-startup/internal/handler/http/campaign"
	"bwa-startup/internal/handler/http/user"
	"bwa-startup/internal/middleware"
	"bwa-startup/internal/repository"
	"bwa-startup/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(r *app.Server, service service.Service, repo repository.Repository, config config.Config) {
	route := r.Group("/")
	route.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("OK")
	})

	api := r.Group("/api/v1")

	userHandler := user.NewHandler(service.UserService(), config.ImageConf(), service.TransactionService())

	authMiddleware := middleware.New(repo.AuthRepository())

	UserRoutePublic(api, userHandler)

	//with middleware
	UserRoutePrivate(api, userHandler, authMiddleware)

	campaignHandler := campaign.NewHandler(service.CampaignService(), service.TransactionService())
	CampaignRoute(api, campaignHandler, authMiddleware)
}
