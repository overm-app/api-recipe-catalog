package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	appErrors "github.com/overm-app/api-recipe-catalog/internal/domain/errors"
	"github.com/overm-app/api-recipe-catalog/internal/domain/models"
	"github.com/overm-app/api-recipe-catalog/internal/domain/ports"
	"github.com/overm-app/api-recipe-catalog/internal/interface/api/response"
	"github.com/overm-app/api-recipe-catalog/internal/usecase"
)

type RecipeHandler struct {
	recipeUseCase *usecase.RecipeUseCase
	jwtService    ports.JWTService
	sugar 	 	  *zap.SugaredLogger
}

func NewRecipeHandler(recipeUseCase *usecase.RecipeUseCase, jwtService ports.JWTService, sugar *zap.SugaredLogger) *RecipeHandler {
	return &RecipeHandler{
		recipeUseCase: recipeUseCase,
		jwtService:    jwtService,
		sugar:         sugar,
	}
}

func (h *RecipeHandler) Create(c *gin.Context) {
	requestID := c.GetString("request_id")
	userID := c.GetString("user_id")
	if userID == "" {
		response.HandleError(c, h.sugar, appErrors.Unauthorized(appErrors.ErrUnauthorized, "User ID missing in context"), requestID)
		return
	}

	var req models.CreateRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(c, h.sugar, appErrors.Validation(appErrors.ErrValidation, err.Error()), requestID)
		return
	}

	resp, err := h.recipeUseCase.Create(c.Request.Context(), userID, &req)
	if err != nil {
		response.HandleError(c, h.sugar, err, requestID)
		return
	}

	c.JSON(201, resp)
}