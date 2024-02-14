package auth

import (
	"bwa-startup/internal/entity"
)

type Repository interface {
	GenerateToken(user *entity.User) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}
