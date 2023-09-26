-- name: CreateMenuItem :one
INSERT INTO "MenuItem" (
  name, description, price, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetMenuItem :one
SELECT * 
FROM "MenuItem" 
WHERE id = $1 AND status = $2
LIMIT 1;

-- name: GetAllMenuItems :many
SELECT *
FROM "MenuItem" 
ORDER BY name 
LIMIT $1 
OFFSET $2;

-- name: GetAllActiveItems :many
SELECT *
FROM "MenuItem" 
WHERE status = 'A'
ORDER BY name 
LIMIT $1 
OFFSET $2;

-- name: UpdateMenuItem :one
UPDATE "MenuItem"
SET name = $1, description = $2, price = $3, updated_at = NOW() 
WHERE id = $4 
RETURNING *;

-- name: DeleteMenuItem :exec
DELETE FROM "MenuItem"
WHERE id = $1;