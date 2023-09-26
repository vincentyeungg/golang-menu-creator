-- name: CreateMenuItemIngredient :one
INSERT INTO "MenuItem_Ingredient" (
    food_id, ingredient_id, status 
)
VALUES ( 
    $1, $2, $3 
)
RETURNING *;

-- name: GetAllIngredientsFromFood :many
SELECT * 
FROM "MenuItem_Ingredient" AS "mii"
JOIN "Ingredient" AS "i" ON "mii".ingredient_id = "i".id
WHERE food_id = $1 
ORDER BY "i".name 
LIMIT $2 
OFFSET $3;

-- name: GetAllActiveIngredientsFromMenu :many
SELECT * 
FROM "MenuItem_Ingredient" AS "mii"
JOIN "Ingredient" AS "i" ON "mii".ingredient_id = "i".id
WHERE food_id = $1 
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
WHERE food_id = $1 AND ingredient_id = $2;