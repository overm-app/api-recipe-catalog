db.recipes.createIndex(
  { user_id: 1, title: 1 },
  { unique: true, partialFilterExpression: { status: "active" }, name: "uniq_active_title_per_user" }
)
db.recipes.createIndex({ user_id: 1, status: 1, created_at: -1 }) // list query