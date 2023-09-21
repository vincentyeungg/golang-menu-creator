-- name: CreateMenuItem :one
INSERT INTO "MenuItem" (
  name, description, price
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetMenuItem :one
SELECT * 
FROM "MenuItem" 
WHERE id = $1 
LIMIT 1;

-- name: GetAllMenuItems :many
SELECT *
FROM "MenuItem" 
LIMIT $1 
OFFSET $2;

-- name: UpdateMenuItem :one
UPDATE "MenuItem"
SET name = $1, description = $2, price = $3 
WHERE id = $4 
RETURNING *;

-- name: DeleteMenuItem :exec
DELETE FROM "MenuItem"
WHERE id = $1;