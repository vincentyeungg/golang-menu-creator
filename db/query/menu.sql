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
WHERE id = $1;

-- name: GetAllMenus :many
SELECT *
FROM "Menu";

-- name: UpdateMenu :one
UPDATE "Menu"
SET name = $1, description = $2
WHERE id = $3 
RETURNING *;

-- name: DeleteMenu :exec
DELETE FROM "Menu"
WHERE id = $1;