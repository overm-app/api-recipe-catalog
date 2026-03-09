const db = db.getSiblingDB("overm_recipes");

db.createCollection("recipes", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      title: "Recipe",
      required: [
        "user_id", "title", "description",
        "steps", "servings", "status",
        "source", "created_at"
      ],
      properties: {
        "_id": { bsonType: "objectId" },
        "user_id": { bsonType: "string" },
        "title": { bsonType: "string" },
        "description": { bsonType: "string" },
        "ingredients": {
          bsonType: "array",
          items: {
            bsonType: "object",
            required: ["name", "quantity", "unit"],
            properties: {
              "name":     { bsonType: "string" },
              "quantity": { bsonType: "double" },
              "unit":     { bsonType: "string" }
            }
          }
        },
        "steps": {
          bsonType: "array",
          items: { bsonType: "string" }
        },
        "servings":  { bsonType: "int" },
        "tags": {
          bsonType: "array",
          items: { bsonType: "string" }
        },
        "macros_per_serving": {
          bsonType: "object",
          properties: {
            "calories":  { bsonType: "double" },
            "protein_g": { bsonType: "double" },
            "carbs_g":   { bsonType: "double" },
            "fat_g":     { bsonType: "double" },
            "fiber_g":   { bsonType: "double" }
          }
        },
        "status":     { bsonType: "string" },
        "source":     { bsonType: "string" },
        "created_at": { bsonType: "date" },
        "updated_at": { bsonType: "date" }
      }
    }
  }
});

db.recipes.createIndex({ "user_id": 1 });
db.recipes.createIndex({ "user_id": 1, "status": 1 });