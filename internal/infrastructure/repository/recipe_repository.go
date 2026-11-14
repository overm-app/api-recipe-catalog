package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	appErrors "github.com/overm-app/api-recipe-catalog/internal/domain/errors"
	"github.com/overm-app/api-recipe-catalog/internal/domain/models"
	"github.com/overm-app/api-recipe-catalog/internal/domain/ports"
	"github.com/overm-app/api-recipe-catalog/internal/infrastructure/db"
)

type recipeRepository struct {
	collection *mongo.Collection
}

func NewRecipeRepository(conn *db.MongoConnection) ports.RecipeRepository {
	return &recipeRepository{
		collection: conn.Database.Collection("recipes"),
	}
}

func (r *recipeRepository) Create(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	if _, err := r.collection.InsertOne(ctx, recipe); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, appErrors.Conflict(appErrors.ErrRecipeAlreadyExists,
				fmt.Sprintf("You already have a recipe titled '%s'", recipe.Title))
		}
		return nil, appErrors.Internal("Failed to create recipe", err)
	}
	return recipe, nil
}

func (r *recipeRepository) GetByID(ctx context.Context, userID, recipeID string) (*models.Recipe, error) {
	var recipe models.Recipe
	err := r.collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: recipeID},
		{Key: "user_id", Value: userID},
		{Key: "status", Value: models.StatusActive},
	}).Decode(&recipe)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, appErrors.NotFound(appErrors.ErrRecipeNotFound,
				fmt.Sprintf("No recipe with id '%s' was found", recipeID))
		}
		return nil, appErrors.Internal("Failed to get recipe", err)
	}
	return &recipe, nil
}

func (r *recipeRepository) GetByUserID(ctx context.Context, userID string, page, pageSize int) ([]models.Recipe, int, error) {
	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "status", Value: models.StatusActive},
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, appErrors.Internal("Failed to count recipes", err)
	}

	opts := options.Find().
		SetSkip(int64((page - 1) * pageSize)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, appErrors.Internal("Failed to list recipes", err)
	}
	defer cursor.Close(ctx)

	recipes := make([]models.Recipe, 0, pageSize)
	if err := cursor.All(ctx, &recipes); err != nil {
		return nil, 0, appErrors.Internal("Failed to decode recipes", err)
	}
	return recipes, int(total), nil
}

func (r *recipeRepository) Update(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	filter := bson.D{
		{Key: "_id", Value: recipe.ID},
		{Key: "user_id", Value: recipe.UserID},
		{Key: "status", Value: models.StatusActive},
	}
	res, err := r.collection.ReplaceOne(ctx, filter, recipe)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) { // renamed to an existing title
			return nil, appErrors.Conflict(appErrors.ErrRecipeAlreadyExists,
				fmt.Sprintf("You already have a recipe titled '%s'", recipe.Title))
		}
		return nil, appErrors.Internal("Failed to update recipe", err)
	}
	if res.MatchedCount == 0 {
		return nil, appErrors.NotFound(appErrors.ErrRecipeNotFound,
			fmt.Sprintf("No recipe with id '%s' was found", recipe.ID))
	}
	return recipe, nil
}

func (r *recipeRepository) Archive(ctx context.Context, userID string, recipeID string) error {
	filter := bson.D{
		{Key: "_id", Value: recipeID},
		{Key: "user_id", Value: userID},
		{Key: "status", Value: models.StatusActive},
	}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "status", Value: models.StatusArchived},
		{Key: "updated_at", Value: time.Now()},
	}}}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return appErrors.Internal("Failed to archive recipe", err)
	}
	if result.MatchedCount == 0 {
		return appErrors.NotFound(appErrors.ErrNotFound, fmt.Sprintf("Recipe %s not found for user %s", recipeID, userID))
	}

	return nil
}
