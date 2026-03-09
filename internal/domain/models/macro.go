package models

type Macro struct {
    Calories float64 `json:"calories"  bson:"calories"`
    ProteinG float64 `json:"protein_g" bson:"protein_g"`
    CarbsG   float64 `json:"carbs_g"   bson:"carbs_g"`
    FatG     float64 `json:"fat_g"     bson:"fat_g"`
    FiberG   float64 `json:"fiber_g"   bson:"fiber_g"`
}