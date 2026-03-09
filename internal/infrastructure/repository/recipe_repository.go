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

func (r *recipeRepository) GetByID(ctx context.Context, userID string, recipeID string) (*models.Recipe, error) {
	var recipe models.Recipe
	err := r.collection.FindOne(ctx, bson.D{
		{Key:"_id", Value: recipeID},
		{Key:"user_id", Value: userID},
		{Key:"status", Value: models.StatusActive},
	}).Decode(&recipe)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} 
		return nil, fmt.Errorf("failed to find recipe by ID: %w", err)
	}
	return &recipe, nil
}

func (r *recipeRepository) GetByUserID(ctx context.Context, userID string, page int, pageSize int) ([]models.Recipe, int, error) {
	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "status", Value: models.StatusActive},
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count recipes: %w", err)
	}

	skip := int64((page - 1) * pageSize)
	opts := options.Find().SetSkip(skip).SetLimit(int64(pageSize)).SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
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
	filter := bson.D{{Key: "_id", Value: recipe.ID}, {Key: "status", Value: models.StatusActive}}

	data, err := bson.Marshal(recipe)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal recipe for update: %w", err)
	}

	updateData := bson.M{}
	if err := bson.Unmarshal(data, &updateData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal recipe for update: %w", err)
	}

	delete(updateData, "_id")
	delete(updateData, "user_id")

	update := bson.D{{Key: "$set", Value: updateData}}
	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update recipe: %w", err)
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
        {Key: "status",     Value: models.StatusArchived},
        {Key: "updated_at", Value: time.Now()},
    }}}

    result, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return fmt.Errorf("failed to archive recipe: %w", err)
    }
	if result.MatchedCount == 0 {
		return fmt.Errorf("recipe not found")
	}

    return nil
}

func (r *recipeRepository) FindByTitle(ctx context.Context, userID string, title string) (*models.Recipe, error) {
	filter := bson.D{
		{Key: "user_id", Value: userID},
		{Key: "title", Value: title},
		{Key: "status", Value: models.StatusActive},
	}

	var recipe models.Recipe
	err := r.collection.FindOne(ctx, filter).Decode(&recipe)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find recipe by title: %w", err)
	}
	return &recipe, nil
}
