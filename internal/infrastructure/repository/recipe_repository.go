package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

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
	_, err := r.collection.InsertOne(ctx, recipe)
	if err != nil {
		return nil, fmt.Errorf("failed to insert recipe: %w", err)
	}
	return recipe, nil
}

func (r *recipeRepository) GetByID(ctx context.Context, id string) (*models.Recipe, error) {
	var recipe models.Recipe
	err := r.collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&recipe)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} 
		return nil, fmt.Errorf("failed to find recipe by ID: %w", err)
	}
	return &recipe, nil
}

func (r *recipeRepository) GetByUserID(ctx context.Context, userID string, page int, pageSize int) ([]models.Recipe, int, error) {
	filer := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "status", Value: models.StatusActive},
	}

	total, err := r.collection.CountDocuments(ctx, filer)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count recipes: %w", err)
	}

	skip := int64((page - 1) * pageSize)
	opts := options.Find().SetSkip(skip).SetLimit(int64(pageSize)).SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, filer, opts)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find recipes: %w", err)
	}
	defer cursor.Close(ctx)

	var recipes []models.Recipe
	if err := cursor.All(ctx, &recipes); err != nil {
		return nil, 0, fmt.Errorf("failed to decode recipes: %w", err)
	}
	return recipes, int(total), nil
}

func (r *recipeRepository) Update(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	recipe.UpdatedAt = time.Now()

	filter := bson.D{{Key: "_id", Value: recipe.ID}}
	update := bson.D{{Key: "$set", Value: recipe}}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update recipe: %w", err)
	}
	return recipe, nil
}

func (r *recipeRepository) Archive(ctx context.Context, id string) error {
    filter := bson.D{{Key: "_id", Value: id}}
    update := bson.D{{Key: "$set", Value: bson.D{
        {Key: "status",     Value: models.StatusArchived},
        {Key: "updated_at", Value: time.Now()},
    }}}

    _, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return fmt.Errorf("failed to archive recipe: %w", err)
    }

    return nil
}
