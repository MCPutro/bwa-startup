package routes

import (
	"bwa-startup/internal/handler/http/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutePublic(group fiber.Router, handler user.Handler) {
	route := group.Group("/user")

	route.Post("/register", handler.RegisterUser)
	route.Post("/login", handler.Login)
	route.Post("/emailCheckers", handler.CheckEmailAvailable)

}

func UserRoutePrivate(group fiber.Router, handler user.Handler, middleware fiber.Handler) {
	route := group.Group("/user")

	route.Post("/uploadAvatar", middleware, handler.UploadAvatar)
	route.Get("/transactions", middleware, handler.FindTrx)
}
