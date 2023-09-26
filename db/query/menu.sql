-- name: CreateMenu :one
INSERT INTO "Menu" (
  name, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetMenu :one
SELECT * 
FROM "Menu" 
WHERE id = $1 
LIMIT 1;

-- name: GetAllMenus :many
SELECT *
FROM "Menu"
ORDER BY name 
LIMIT $1 
OFFSET $2;

-- name: GetAllActiveMenus :many
SELECT *
FROM "Menu"
WHERE status = 'A'
ORDER BY name 
LIMIT $1 
OFFSET $2;

-- name: UpdateMenu :one
UPDATE "Menu"
SET name = $1, description = $2, updated_at = NOW() 
WHERE id = $3 
RETURNING *;

-- name: DeleteMenu :exec
DELETE FROM "Menu"
WHERE id = $1;