package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

var refreshTokenTTL time.Duration

func InitRefreshTokenTTL(days int) {
	if days > 0 {
		refreshTokenTTL = time.Duration(days) * 24 * time.Hour
	} else {
		refreshTokenTTL = 30 * 24 * time.Hour
	}
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 64)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
