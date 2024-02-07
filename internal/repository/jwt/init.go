package jwt

import (
	"bwa-startup/config"
	"bwa-startup/internal/entity"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type crypto struct {
	auth config.AuthConfig
}

//const JWT_SERCRET_KEY = "my_secret"

func (j *crypto) GenerateJWT(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"Id":    user.ID,
		"Email": user.Email,
		//"iss":      "issue",
		"sub": "login",
		//"aud": "aud",
		"exp": jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		//"nbf" : jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), //NotBefore
		"jti": uuid.NewString(), //unique id jwt
	})

	ss, err := token.SignedString([]byte(j.auth.SecretKey()))

	if err != nil {
		return "", err
	}
	return ss, nil
}

func (j *crypto) ValidateJWT(token string) (map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.auth.SecretKey()), nil
	})
	if err != nil {
		return nil, err
	}

	if t.Valid {
		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			return map[string]interface{}{}, errors.New("token validation error")
		}
		payload := make(map[string]interface{})

		for key, value := range claims {
			payload[key] = value
		}

		return payload, nil
	} else {
		return nil, jwtError(err)
	}
}

func jwtError(err error) error {
	if errors.Is(err, jwt.ErrInvalidKey) {
		return jwt.ErrInvalidKey
	} else if errors.Is(err, jwt.ErrInvalidKeyType) {
		return jwt.ErrInvalidKeyType
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return jwt.ErrTokenMalformed
	} else if errors.Is(err, jwt.ErrTokenUnverifiable) {
		return jwt.ErrTokenUnverifiable
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return jwt.ErrTokenSignatureInvalid
	} else if errors.Is(err, jwt.ErrTokenRequiredClaimMissing) {
		return jwt.ErrTokenRequiredClaimMissing
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return jwt.ErrTokenExpired
	} else if errors.Is(err, jwt.ErrTokenInvalidSubject) {
		return jwt.ErrTokenInvalidSubject
	} else if errors.Is(err, jwt.ErrTokenInvalidId) {
		return jwt.ErrTokenInvalidId
	} else if errors.Is(err, jwt.ErrTokenInvalidClaims) {
		return jwt.ErrTokenInvalidClaims
	} else if errors.Is(err, jwt.ErrInvalidType) {
		return jwt.ErrInvalidType
	} else {
		return err
	}
}

func NewJWT(auth config.AuthConfig) Repository {
	return &crypto{auth: auth}
}
