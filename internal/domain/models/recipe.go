package models

import (
	"time"
)

type Recipe struct {
    ID               string       `json:"id"                bson:"_id"`
    UserID           string       `json:"-"                 bson:"user_id"`
    Title            string       `json:"title"             bson:"title"`
    Description      string       `json:"description"       bson:"description"`
    Ingredients      []Ingredient `json:"ingredients"       bson:"ingredients"`
    Steps            []string     `json:"steps"             bson:"steps"`
    Servings         int          `json:"servings"          bson:"servings"`
    Tags             []string     `json:"tags"              bson:"tags"`
    MacrosPerServing Macro        `json:"macros_per_serving" bson:"macros_per_serving"`
    Status           string       `json:"status"            bson:"status"`
    Source           string       `json:"source"            bson:"source"`
    CreatedAt        time.Time    `json:"created_at"        bson:"created_at"`
    UpdatedAt        time.Time    `json:"updated_at"        bson:"updated_at"`
}

type RecipeListResponse struct {
    Data []Recipe `json:"data"`
    Meta Meta     `json:"meta"`
}

type CreateRecipeRequest struct {
	Title       string       `json:"title" binding:"required"`
	Description string       `json:"description" binding:"required"`
	Ingredients []Ingredient `json:"ingredients" binding:"required,dive"`
	Steps       []string     `json:"steps" binding:"required,dive,gt=0"`
	Servings    int         `json:"servings" binding:"required,gt=0"`
	Tags        []string     `json:"tags" binding:"dive,gt=0"`
}

type UpdateRecipeRequest struct {
	Title       *string       `json:"title,omitempty"`
	Description *string       `json:"description,omitempty"`
	Ingredients *[]Ingredient `json:"ingredients,omitempty"`
	Steps       *[]string     `json:"steps,omitempty"`
	Servings    *int         `json:"servings,omitempty"`
	Tags        *[]string     `json:"tags,omitempty"`
}

type Meta struct {
    Total    int `json:"total"`
    Page     int `json:"page"`
    PageSize int `json:"page_size"`
}

const (
	StatusActive      = "active"
	StatusArchived    = "archived"
	SourceManual      = "manual"
	SourceTranscribed = "transcribed"
)
