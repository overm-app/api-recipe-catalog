package service

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"

	appErrors "github.com/overm-app/api-recipe-catalog/internal/domain/errors"
	"github.com/overm-app/api-recipe-catalog/internal/domain/models"
	"github.com/overm-app/api-recipe-catalog/internal/domain/ports"
)

type JWTService struct {
	secret []byte
}

func NewJWTService(secret []byte) ports.JWTService {
	return &JWTService{
		secret: secret,
	}
}

func (j *JWTService) ValidateToken(tokenString string) (*models.JWTClaims, error) {
	parsed, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, appErrors.Unauthorized(appErrors.ErrUnauthorized, "Unexpected signing method")
		}
		return j.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, appErrors.Unauthorized(appErrors.ErrTokenExpired, "Token has expired")
		}
		return nil, appErrors.Unauthorized(appErrors.ErrUnauthorized, "Invalid token: "+err.Error())
	}

	claims, ok := parsed.Claims.(*models.JWTClaims)
	if !ok || !parsed.Valid {
		return nil, appErrors.Unauthorized(appErrors.ErrUnauthorized, "Invalid token claims")
	}

	return claims, nil
}
