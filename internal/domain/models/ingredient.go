package models

type Ingredient struct {
    Name     string  `json:"name"     bson:"name"`
    Quantity float64 `json:"quantity" bson:"quantity"`
    Unit     string  `json:"unit"     bson:"unit"`
}