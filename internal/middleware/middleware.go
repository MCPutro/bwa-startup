package middleware

import (
	"bwa-startup/internal/handler/response"
	"bwa-startup/internal/repository/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func New(jwtRepository jwt.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.JSON(http.StatusUnauthorized, response.New{
				Success:      false,
				Code:         http.StatusUnauthorized,
				ErrorMessage: "Authorization header is required",
			})

			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token2, err := jwtRepository.ValidateJWT(tokenString)
		if err != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.JSON(http.StatusUnauthorized, response.New{
				Success:      false,
				Code:         http.StatusUnauthorized,
				ErrorMessage: err.Error(),
			})
			c.Abort()
			return
		}

		userID, ok := token2["Id"]
		if !ok {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
			c.JSON(http.StatusUnauthorized, response.New{
				Success:      false,
				Code:         http.StatusUnauthorized,
				ErrorMessage: "Invalid JWT token",
			})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
