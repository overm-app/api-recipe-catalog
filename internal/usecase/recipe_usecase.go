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
		Ingredients:      req.Ingredients,
		Steps:            req.Steps,
		Servings:         req.Servings,
		Tags:             req.Tags,
		MacrosPerServing: models.Macro{},
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

func (uc *RecipeUseCase) ListOne(ctx context.Context, userID, recipeID string) (*models.Recipe, error) {
	recipe, err := uc.recipeRepo.GetByID(ctx, userID, recipeID)
	if err != nil {
		return nil, appErrors.Internal("Failed to retrieve recipe", err)
	}
	if recipe == nil {
		return nil, appErrors.NotFound(appErrors.ErrNotFound, "Recipe not found")
	}
	return recipe, nil
}

func (uc *RecipeUseCase) List(ctx context.Context, userID string, page int, pageSize int) (*models.RecipeListResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }

    recipes, total, err := uc.recipeRepo.GetByUserID(ctx, userID, page, pageSize)
    if err != nil {
        return nil, appErrors.Internal("Failed to retrieve recipes", err)
    }

    return &models.RecipeListResponse{
        Data: recipes,
        Meta: models.Meta{
            Total:    total,
            Page:     page,
            PageSize: pageSize,
        },
    }, nil
}

func (uc *RecipeUseCase) Update(ctx context.Context, userID string, id string, req *models.UpdateRecipeRequest) (*models.Recipe, error) {
    recipe, err := uc.recipeRepo.GetByID(ctx, id, userID)
    if err != nil {
        return nil, appErrors.Internal("Failed to retrieve recipe", err)
    }
    if recipe == nil {
        return nil, appErrors.NotFound(appErrors.ErrNotFound, "Recipe not found")
    }

    if req.Title != nil && *req.Title != recipe.Title {
        existing, err := uc.recipeRepo.FindByTitle(ctx, userID, *req.Title)
        if err != nil {
            return nil, appErrors.Internal("Failed to check existing recipes", err)
        }
        if existing != nil {
            return nil, appErrors.Conflict(appErrors.ErrRecipeAlreadyExists, "A recipe with this title already exists")
        }
    }

    if req.Title != nil {
        recipe.Title = *req.Title
    }
    if req.Description != nil {
        recipe.Description = *req.Description
    }
    if req.Ingredients != nil && len(*req.Ingredients) > 0 {
        recipe.Ingredients = *req.Ingredients
    }
    if req.Steps != nil && len(*req.Steps) > 0 {
        recipe.Steps = *req.Steps
    }
    if req.Servings != nil {
        recipe.Servings = *req.Servings
    }
    if req.Tags != nil {
        recipe.Tags = *req.Tags
    }

    recipe.UpdatedAt = time.Now()

    updated, err := uc.recipeRepo.Update(ctx, recipe)
    if err != nil {
        return nil, appErrors.Internal("Failed to update recipe", err)
    }

    return updated, nil
}

func (uc *RecipeUseCase) Archive(ctx context.Context, userID string, id string) error {
    err := uc.recipeRepo.Archive(ctx, id, userID)
    if err != nil {
        if err.Error() == "recipe not found" {
            return appErrors.NotFound(appErrors.ErrNotFound, "Recipe not found")
        }
        return appErrors.Internal("Failed to archive recipe", err)
    }
    return nil
}