package ports

import (
	"context"

	"github.com/overm-app/api-recipe-catalog/internal/domain/models"
)

type RecipeRepository interface {
	Create(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error)
	GetByID(ctx context.Context, id string) (*models.Recipe, error)
	GetByUserID(ctx context.Context, userID string, page int, pageSize int) ([]models.Recipe, int, error)
	Update(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error)
	Archive(ctx context.Context, id string) error
	FindByTitle(ctx context.Context, userID string, title string) (*models.Recipe, error)
}
