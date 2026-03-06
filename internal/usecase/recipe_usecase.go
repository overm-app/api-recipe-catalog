package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	appErrors "github.com/overm-app/api-recipe-catalog/internal/domain/errors"
	"github.com/overm-app/api-recipe-catalog/internal/domain/models"
	"github.com/overm-app/api-recipe-catalog/internal/domain/ports"
)

type RecipeUseCase struct {
	recipeRepo ports.RecipeRepository
}

func NewRecipeUseCase(recipeRepo ports.RecipeRepository) *RecipeUseCase {
	return &RecipeUseCase{
		recipeRepo: recipeRepo,
	}
}

func (uc *RecipeUseCase) Create(ctx context.Context, userID string, req *models.CreateRecipeRequest) (*models.Recipe, error) {
	existing, err := uc.recipeRepo.FindByTitle(ctx, userID, req.Title)
	if err != nil {
		return nil, appErrors.Internal("Failed to check existing recipes", err)
	}
	if existing != nil {
		return nil, appErrors.Conflict(appErrors.ErrRecipeAlreadyExists, "A recipe with this title already exists")
	}

	recipe := &models.Recipe{
		ID:               uuid.New().String(),
		UserID:           userID,
		Title:            req.Title,
		Description:      req.Description,
		Ingredients:      req.Ingredients, // stored as-is, no calculation
		Steps:            req.Steps,
		Servings:         req.Servings,
		Tags:             req.Tags,
		MacrosPerServing: models.Macro{}, // empty placeholder
		Status:           models.StatusActive,
		Source:           models.SourceManual,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	recipe, err = uc.recipeRepo.Create(ctx, recipe)
	if err != nil {
		return nil, appErrors.Internal("Failed to create recipe", err)
	}

	return recipe, nil
}
