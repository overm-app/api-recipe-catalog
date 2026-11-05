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
	sugar         *zap.SugaredLogger
}

func NewRecipeHandler(recipeUseCase *usecase.RecipeUseCase, jwtService ports.JWTService, sugar *zap.SugaredLogger) *RecipeHandler {
	return &RecipeHandler{
		recipeUseCase: recipeUseCase,
		jwtService:    jwtService,
		sugar:         sugar,
	}
}

func (h *RecipeHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		response.HandleError(c, appErrors.Unauthorized(appErrors.ErrUnauthorized, "User ID missing in context"))
		return
	}

	var req models.CreateRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(c, appErrors.Validation(appErrors.ErrValidation, err.Error()))
		return
	}

	resp, err := h.recipeUseCase.Create(c.Request.Context(), userID, &req)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.JSON(201, resp)
}

func (h *RecipeHandler) GetByID(c *gin.Context) {
	userID := c.GetString("user_id")
	recipeID := c.Param("id")

	resp, err := h.recipeUseCase.GetByID(c.Request.Context(), userID, recipeID)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(200, resp)
}

func (h *RecipeHandler) List(c *gin.Context) {
	userID := c.GetString("user_id")

	page, pageSize := response.GetPaginationParams(c)

	resp, err := h.recipeUseCase.List(c.Request.Context(), userID, page, pageSize)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(200, resp)
}

func (h *RecipeHandler) Update(c *gin.Context) {
	userID := c.GetString("user_id")
	recipeID := c.Param("id")

	var req models.UpdateRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleError(c, appErrors.Validation(appErrors.ErrValidation, err.Error()))
		return
	}

	resp, err := h.recipeUseCase.Update(c.Request.Context(), userID, recipeID, &req)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	c.JSON(200, resp)
}

func (h *RecipeHandler) Archive(c *gin.Context) {
	userID := c.GetString("user_id")
	recipeID := c.Param("id")

	if err := h.recipeUseCase.Archive(c.Request.Context(), userID, recipeID); err != nil {
		response.HandleError(c, err)
		return
	}
	c.Status(204)
}
