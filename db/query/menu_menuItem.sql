-- name: getAllItemsFromMenu :many
SELECT * 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 
LIMIT $2 
OFFSET $3;