package user

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	RegisterUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	CheckEmailAvailable(c *fiber.Ctx) error
	UploadAvatar(c *fiber.Ctx) error
	FindTrx(c *fiber.Ctx) error
}
