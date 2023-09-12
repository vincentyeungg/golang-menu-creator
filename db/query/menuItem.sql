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
WHERE id = $1;

-- name: GetAllMenuItems :many
SELECT *
FROM "MenuItem";

-- name: UpdateMenuItem :one
UPDATE "MenuItem"
SET name = $1, description = $2, price = $3 
WHERE id = $4 
RETURNING *;

-- name: DeleteMenuItem :exec
DELETE FROM "MenuItem"
WHERE id = $1;