package service

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"

	"github.com/overm-app/api-recipe-catalog/internal/domain/models"
	"github.com/overm-app/api-recipe-catalog/internal/domain/ports"
)

type JWTService struct {
	secret     []byte
}

func NewJWTService(secret []byte) ports.JWTService {
	return &JWTService{
		secret:     secret,
	}
}

func (j *JWTService) ValidateToken(tokenString string) (*models.JWTClaims, error) {
	parsed, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return j.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to parse JWT: %w", err)
	}

	claims, ok := parsed.Claims.(*models.JWTClaims)
	if !ok || !parsed.Valid {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil
}
