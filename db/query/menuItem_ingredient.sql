-- name: CreateMenuItemIngredient :one
INSERT INTO "MenuItem_Ingredient" (
    food_id, ingredient_id, status, created_by, updated_by
)
VALUES ( 
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetMenuItemIngredient :one
SELECT * 
FROM "MenuItem_Ingredient" 
WHERE id = $1 
LIMIT 1;

-- name: GetAllIngredientsFromFood :many
SELECT * 
FROM "MenuItem_Ingredient" AS "mii"
JOIN "Ingredient" AS "i" ON "mii".ingredient_id = "i".id
WHERE food_id = $1 
ORDER BY "i".name 
LIMIT $2 
OFFSET $3;

-- name: GetAllActiveIngredientsFromFood :many
SELECT * 
FROM "MenuItem_Ingredient" AS "mii"
JOIN "Ingredient" AS "i" ON "mii".ingredient_id = "i".id
WHERE food_id = $1 AND "mii".status = 'A' 
ORDER BY "i".name 
LIMIT $2 
OFFSET $3;

-- name: GetActiveIngredientFromMenu :one
SELECT * 
FROM "MenuItem_Ingredient" AS "mii"
JOIN "Ingredient" AS "i" ON "mii".ingredient_id = "i".id
WHERE food_id = $1 
ORDER BY "i".name 
LIMIT $2 
OFFSET $3;

-- name: DeleteIngredientFromItem :exec
DELETE FROM "MenuItem_Ingredient" 
WHERE food_id = $1 AND ingredient_id = $2 AND created_by = $3 ;