package ports

import "github.com/overm-app/api-recipe-catalog/internal/domain/models"

type JWTService interface {
    ValidateToken(tokenString string) (*models.JWTClaims, error)
}