package middleware

import (
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/auth"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func New(jwtRepository auth.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get(fiber.HeaderAuthorization)

		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(
				response.New{
					Success: false,
					Code:    http.StatusUnauthorized,
					Message: "authorization header is required",
				})
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token2, err := jwtRepository.ValidateToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(response.New{
				Success: false,
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
		}

		userID, ok := token2["Id"]
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(response.New{
				Success: false,
				Code:    http.StatusUnauthorized,
				Message: "invalid token",
			})
		}

		c.Locals("userID", userID)

		return c.Next()
	}
}
