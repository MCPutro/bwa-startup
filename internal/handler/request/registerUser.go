package request

import (
	"bwa-startup/internal/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (r *RegisterUser) ToEntity() *entity.User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	return &entity.User{
		Name:       r.Name,
		Occupation: r.Occupation,
		Email:      r.Email,
		Password:   string(passwordHash),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
