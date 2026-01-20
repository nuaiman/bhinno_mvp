package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte
var accessTokenTTL = 15 * time.Minute

func InitJWT(key string, ttlMinutes int) {
	jwtKey = []byte(key)
	if ttlMinutes > 0 {
		accessTokenTTL = time.Duration(ttlMinutes) * time.Minute
	}
}

type CustomClaims struct {
	UserID int64  `json:"user_id"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int64, phone, role string) (string, error) {
	if len(jwtKey) == 0 {
		return "", errors.New("jwt key not initialized")
	}

	claims := &CustomClaims{
		UserID: userID,
		Phone:  phone,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   phone,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyJWT(tokenStr string) (userID int64, phone string, role string, err error) {
	if len(jwtKey) == 0 {
		return 0, "", "", errors.New("jwt key not initialized")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, "", "", err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, "", "", errors.New("invalid claims type")
	}

	if claims.UserID == 0 || claims.Phone == "" {
		return 0, "", "", errors.New("invalid user claims")
	}

	return claims.UserID, claims.Phone, claims.Role, nil
}
