package config

import "time"

type AuthConfig interface {
	SecretKey() string
	TokenExpiredTime() time.Time
}

type JWT struct {
	TokenSecretKey   string `mapstructure:"secret-key"`
	TokenExpiredTIme int    `mapstructure:"expired-time-in-minute"`
}

func (J *JWT) SecretKey() string {
	return J.TokenSecretKey
}

func (J *JWT) TokenExpiredTime() time.Time {
	t := time.Duration(J.TokenExpiredTIme)
	return time.Now().UTC().Add(t * time.Minute)
}
