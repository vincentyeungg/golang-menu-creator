-- name: CreateMenuItem :one
INSERT INTO "MenuItem" (
  name, description, price
) VALUES (
  $1, $2, $3
)
RETURNING *;