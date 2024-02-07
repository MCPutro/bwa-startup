package jwt

import (
	"bwa-startup/internal/entity"
)

type Repository interface {
	GenerateJWT(user *entity.User) (string, error)
	ValidateJWT(token string) (map[string]interface{}, error)
}
