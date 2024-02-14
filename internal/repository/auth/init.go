package auth

import (
	"bwa-startup/config"
	"bwa-startup/internal/entity"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type crypto struct {
	auth config.AuthConfig
}

func (j *crypto) GenerateToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"Id":    user.ID,
		"Email": user.Email,
		//"iss":      "issue",
		"sub": "login",
		//"aud": "aud",
		//"exp": jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		"exp": jwt.NewNumericDate(j.auth.TokenExpiredTime()),
		//"nbf" : jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), //NotBefore
		"jti": uuid.NewString(), //unique id jwt
	})

	tokenString, err := token.SignedString([]byte(j.auth.SecretKey()))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *crypto) ValidateToken(token string) (map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			//return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
			return nil, errors.New("invalid token")
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

func NewAuth(auth config.AuthConfig) Repository {
	return &crypto{auth: auth}
}
