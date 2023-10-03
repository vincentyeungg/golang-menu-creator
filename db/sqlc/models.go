// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Ingredient struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	BrandName   string    `json:"brand_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	Status      string    `json:"status"`
}

type Menu struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	Status      string    `json:"status"`
}

type MenuItem struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
	Status      string    `json:"status"`
}

type MenuItemIngredient struct {
	ID           int32     `json:"id"`
	FoodID       int32     `json:"food_id"`
	IngredientID int32     `json:"ingredient_id"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
	Status       string    `json:"status"`
}

type MenuMenuItem struct {
	ID        int32     `json:"id"`
	MenuID    int32     `json:"menu_id"`
	FoodID    int32     `json:"food_id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	Status    string    `json:"status"`
}