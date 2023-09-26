-- name: CreateMenuMenuItem :one
INSERT INTO "Menu_MenuItem" (
    menu_id, food_id, status 
)
VALUES ( 
    $1, $2, $3 
)
RETURNING *;

-- name: GetAllItemsFromMenu :many
SELECT "mmi".menu_id, "mi".name, "mi".description, "mi".price 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 
ORDER BY "mi".name 
LIMIT $2 
OFFSET $3;

-- name: GetAllActiveItemsFromMenu :many
SELECT "mmi".menu_id, "mi".name, "mi".description, "mi".price 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 AND status = 'A' 
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

-- name: DeleteMenuFromMenu :exec
DELETE FROM "Menu_MenuItem" 
WHERE menu_id = $1 AND food_id = $2;