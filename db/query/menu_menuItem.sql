-- name: CreateMenuMenuItem :one
INSERT INTO "Menu_MenuItem" (
    menu_id, food_id, status, created_by, updated_by
)
VALUES ( 
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetMenuMenuItem :one
SELECT * 
FROM "Menu_MenuItem" 
WHERE id = $1 
LIMIT 1;

-- name: GetAllItemsFromMenu :many
SELECT "mmi".menu_id, "mi".name, "mi".description, "mi".price, "mi".status 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 
ORDER BY "mi".name 
LIMIT $2 
OFFSET $3;

-- name: GetAllActiveItemsFromMenu :many
SELECT "mmi".menu_id, "mi".name, "mi".description, "mi".price, "mi".status 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 AND "mmi".status = 'A' 
ORDER BY "mi".name 
LIMIT $2 
OFFSET $3;

-- name: GetActiveItemFromMenu :one
SELECT "mmi".menu_id, "mi".name, "mi".description, "mi".price 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 AND food_id = $2 AND status = 'A' 
ORDER BY "mi".name 
LIMIT $3 
OFFSET $4;

-- name: DeleteMenuItemFromMenu :exec
DELETE FROM "Menu_MenuItem" 
WHERE menu_id = $1 AND food_id = $2 AND created_by = $3;