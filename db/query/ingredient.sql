-- name: CreateIngredient :one
INSERT INTO "Ingredient" (
  name, brand_name, description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetIngredient :one
SELECT * 
FROM "Ingredient" 
WHERE id = $1 
LIMIT 1;

-- name: GetAllIngredient :many
SELECT *
FROM "Ingredient" 
ORDER BY name 
LIMIT $1 
OFFSET $2;

-- name: GetAllActiveIngredients :many
SELECT *
FROM "Ingredient" 
WHERE status = 'A'
ORDER BY name 
LIMIT $1 
OFFSET $2;

-- name: UpdateIngredient :one
UPDATE "Ingredient"
SET name = $1, brand_name = $2, description = $3, updated_at = NOW()
WHERE id = $4 
RETURNING *;

-- name: DeleteIngredient :exec
DELETE FROM "Ingredient"
WHERE id = $1;