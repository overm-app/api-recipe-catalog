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

const (
	defaultPageSize = 20
	maxPageSize     = 100
)

func (uc *RecipeUseCase) Create(ctx context.Context, userID string, req *models.CreateRecipeRequest) (*models.Recipe, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, appErrors.Internal("Failed to generate recipe id", err)
	}
	now := time.Now().UTC()

	recipe := &models.Recipe{
		ID:          id.String(),
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Ingredients: req.Ingredients,
		Steps:       req.Steps,
		Servings:    req.Servings,
		Tags:        req.Tags,
		Status:      models.StatusActive,
		Source:      models.SourceManual,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	// TODO: nutrition call here; on failure leave MacrosPerServing nil and continue.
	return uc.recipeRepo.Create(ctx, recipe)
}

func (uc *RecipeUseCase) GetByID(ctx context.Context, userID, recipeID string) (*models.Recipe, error) {
	return uc.recipeRepo.GetByID(ctx, userID, recipeID)
}

func (uc *RecipeUseCase) List(ctx context.Context, userID string, page, pageSize int) (*models.RecipeListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	recipes, total, err := uc.recipeRepo.GetByUserID(ctx, userID, page, pageSize)
	if err != nil {
		return nil, err
	}

	return &models.RecipeListResponse{
		Data: recipes,
		Meta: models.Meta{
			Total:      total,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: (total + pageSize - 1) / pageSize,
		},
	}, nil
}

func (uc *RecipeUseCase) Update(ctx context.Context, userID, recipeID string, req *models.UpdateRecipeRequest) (*models.Recipe, error) {
	recipe, err := uc.recipeRepo.GetByID(ctx, userID, recipeID)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		recipe.Title = *req.Title
	}
	if req.Description != nil {
		recipe.Description = *req.Description
	}
	if req.Ingredients != nil {
		recipe.Ingredients = *req.Ingredients
	}
	if req.Steps != nil {
		recipe.Steps = *req.Steps
	}
	if req.Servings != nil {
		recipe.Servings = *req.Servings
	}
	if req.Tags != nil {
		recipe.Tags = *req.Tags
	}
	recipe.UpdatedAt = time.Now().UTC()

	// TODO: if req.Ingredients != nil || req.Servings != nil, recalculate macros.
	return uc.recipeRepo.Update(ctx, recipe)
}

func (uc *RecipeUseCase) Archive(ctx context.Context, userID, recipeID string) error {
	return uc.recipeRepo.Archive(ctx, userID, recipeID)
}
