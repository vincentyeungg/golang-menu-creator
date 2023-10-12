// Code generated by sqlc. DO NOT EDIT.
// source: menu_menuItem.sql

package db

import (
	"context"
)

const createMenuMenuItem = `-- name: CreateMenuMenuItem :one
INSERT INTO "Menu_MenuItem" (
    menu_id, food_id, status, created_by, updated_by
)
VALUES ( 
    $1, $2, $3, $4, $5
)
RETURNING id, menu_id, food_id, created_at, created_by, updated_at, updated_by, status
`

type CreateMenuMenuItemParams struct {
	MenuID    int32  `json:"menu_id"`
	FoodID    int32  `json:"food_id"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

func (q *Queries) CreateMenuMenuItem(ctx context.Context, arg CreateMenuMenuItemParams) (MenuMenuItem, error) {
	row := q.db.QueryRowContext(ctx, createMenuMenuItem,
		arg.MenuID,
		arg.FoodID,
		arg.Status,
		arg.CreatedBy,
		arg.UpdatedBy,
	)
	var i MenuMenuItem
	err := row.Scan(
		&i.ID,
		&i.MenuID,
		&i.FoodID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Status,
	)
	return i, err
}

const deleteMenuItemFromMenu = `-- name: DeleteMenuItemFromMenu :exec
DELETE FROM "Menu_MenuItem" 
WHERE menu_id = $1 AND food_id = $2 AND created_by = $3
`

type DeleteMenuItemFromMenuParams struct {
	MenuID    int32  `json:"menu_id"`
	FoodID    int32  `json:"food_id"`
	CreatedBy string `json:"created_by"`
}

func (q *Queries) DeleteMenuItemFromMenu(ctx context.Context, arg DeleteMenuItemFromMenuParams) error {
	_, err := q.db.ExecContext(ctx, deleteMenuItemFromMenu, arg.MenuID, arg.FoodID, arg.CreatedBy)
	return err
}

const getActiveItemFromMenu = `-- name: GetActiveItemFromMenu :one
SELECT "mmi".menu_id, "mi".name, "mi".description, "mi".price 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 AND food_id = $2 AND status = 'A' 
ORDER BY "mi".name 
LIMIT $3 
OFFSET $4
`

type GetActiveItemFromMenuParams struct {
	MenuID int32 `json:"menu_id"`
	FoodID int32 `json:"food_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetActiveItemFromMenuRow struct {
	MenuID      int32  `json:"menu_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

func (q *Queries) GetActiveItemFromMenu(ctx context.Context, arg GetActiveItemFromMenuParams) (GetActiveItemFromMenuRow, error) {
	row := q.db.QueryRowContext(ctx, getActiveItemFromMenu,
		arg.MenuID,
		arg.FoodID,
		arg.Limit,
		arg.Offset,
	)
	var i GetActiveItemFromMenuRow
	err := row.Scan(
		&i.MenuID,
		&i.Name,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const getAllActiveItemsFromMenu = `-- name: GetAllActiveItemsFromMenu :many
SELECT "mmi".menu_id, "mi".name, "mi".description, "mi".price, "mi".status 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 AND "mmi".status = 'A' 
ORDER BY "mi".name 
LIMIT $2 
OFFSET $3
`

type GetAllActiveItemsFromMenuParams struct {
	MenuID int32 `json:"menu_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAllActiveItemsFromMenuRow struct {
	MenuID      int32  `json:"menu_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Status      string `json:"status"`
}

func (q *Queries) GetAllActiveItemsFromMenu(ctx context.Context, arg GetAllActiveItemsFromMenuParams) ([]GetAllActiveItemsFromMenuRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllActiveItemsFromMenu, arg.MenuID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllActiveItemsFromMenuRow{}
	for rows.Next() {
		var i GetAllActiveItemsFromMenuRow
		if err := rows.Scan(
			&i.MenuID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllItemsFromMenu = `-- name: GetAllItemsFromMenu :many
SELECT "mmi".menu_id, "mi".name, "mi".description, "mi".price, "mi".status 
FROM "Menu_MenuItem" AS "mmi"
JOIN "MenuItem" AS "mi" ON "mmi".food_id = "mi".id
WHERE menu_id = $1 
ORDER BY "mi".name 
LIMIT $2 
OFFSET $3
`

type GetAllItemsFromMenuParams struct {
	MenuID int32 `json:"menu_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetAllItemsFromMenuRow struct {
	MenuID      int32  `json:"menu_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Status      string `json:"status"`
}

func (q *Queries) GetAllItemsFromMenu(ctx context.Context, arg GetAllItemsFromMenuParams) ([]GetAllItemsFromMenuRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllItemsFromMenu, arg.MenuID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllItemsFromMenuRow{}
	for rows.Next() {
		var i GetAllItemsFromMenuRow
		if err := rows.Scan(
			&i.MenuID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMenuMenuItem = `-- name: GetMenuMenuItem :one
SELECT id, menu_id, food_id, created_at, created_by, updated_at, updated_by, status 
FROM "Menu_MenuItem" 
WHERE id = $1 
LIMIT 1
`

func (q *Queries) GetMenuMenuItem(ctx context.Context, id int32) (MenuMenuItem, error) {
	row := q.db.QueryRowContext(ctx, getMenuMenuItem, id)
	var i MenuMenuItem
	err := row.Scan(
		&i.ID,
		&i.MenuID,
		&i.FoodID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.Status,
	)
	return i, err
}
