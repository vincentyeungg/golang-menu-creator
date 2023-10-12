package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomMenuItemIngredient(t *testing.T) MenuItemIngredient {
	// create menu item and ingredient
	menuItem := createRandomMenuItem(t)
	ingredient := createRandomIngredient(t)

	arg := CreateMenuItemIngredientParams{
		FoodID:       menuItem.ID,
		IngredientID: ingredient.ID,
		Status:       "A",
		CreatedBy:    menuItem.CreatedBy,
	}

	res, err := testQueries.CreateMenuItemIngredient(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, arg.FoodID, res.FoodID)
	require.Equal(t, arg.IngredientID, res.IngredientID)
	require.Equal(t, arg.Status, res.Status)
	require.Equal(t, arg.CreatedBy, res.CreatedBy)

	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)

	return res
}

func TestCreateRandomMenuItemIngredient(t *testing.T) {
	CreateRandomMenuItemIngredient(t)
}

func TestGetMenuItemIngredient(t *testing.T) {
	menuItemIngredient := CreateRandomMenuItemIngredient(t)

	res, err := testQueries.GetMenuItemIngredient(context.Background(), menuItemIngredient.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, menuItemIngredient.ID, res.ID)
	require.Equal(t, menuItemIngredient.FoodID, res.FoodID)
	require.Equal(t, menuItemIngredient.FoodID, res.FoodID)
	require.Equal(t, menuItemIngredient.CreatedBy, res.CreatedBy)
	require.Equal(t, menuItemIngredient.UpdatedBy, res.UpdatedBy)
	require.Equal(t, menuItemIngredient.Status, res.Status)
}

func TestGetActiveIngredientFromFood(t *testing.T) {
	menuItemIngredient := CreateRandomMenuItemIngredient(t)

	res, err := testQueries.GetMenuItemIngredient(context.Background(), menuItemIngredient.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, menuItemIngredient.ID, res.ID)
	require.Equal(t, menuItemIngredient.FoodID, res.FoodID)
	require.Equal(t, menuItemIngredient.FoodID, res.FoodID)
	require.Equal(t, menuItemIngredient.CreatedBy, res.CreatedBy)
	require.Equal(t, menuItemIngredient.UpdatedBy, res.UpdatedBy)
	require.Equal(t, res.Status, "A")
}

func TestDeleteMenuItemIngredient(t *testing.T) {
	menuItemIngredient := CreateRandomMenuItemIngredient(t)

	arg := DeleteIngredientFromItemParams{
		FoodID:       menuItemIngredient.FoodID,
		IngredientID: menuItemIngredient.IngredientID,
		CreatedBy:    menuItemIngredient.CreatedBy,
	}

	err := testQueries.DeleteIngredientFromItem(context.Background(), arg)
	res, err := testQueries.GetMenuItemIngredient(context.Background(), menuItemIngredient.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}

func TestGetAllIngredientsFromFood(t *testing.T) {
	menuItem := createRandomMenuItem(t)
	menuItemIngredients := []MenuItemIngredient{}

	for i := 0; i < 10; i++ {
		ingredient := createRandomIngredient(t)
		arg := CreateMenuItemIngredientParams{
			FoodID:       menuItem.ID,
			IngredientID: ingredient.ID,
			Status:       "A",
			CreatedBy:    menuItem.CreatedBy,
			UpdatedBy:    menuItem.UpdatedBy,
		}

		res, err := testQueries.CreateMenuItemIngredient(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		menuItemIngredients = append(menuItemIngredients, res)
	}

	arg := GetAllIngredientsFromFoodParams{
		FoodID: menuItem.ID,
		Limit:  5,
		Offset: 5,
	}

	res, err := testQueries.GetAllIngredientsFromFood(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, res, 5)

	for _, item := range res {
		require.NotEmpty(t, item)
		require.Equal(t, item.FoodID, menuItem.ID)
		require.Equal(t, item.CreatedBy, menuItem.CreatedBy)
		require.Equal(t, item.Status, menuItem.Status)
	}
}

func TestGetAllActiveIngredientsFromFood(t *testing.T) {
	menuItem := createRandomMenuItem(t)
	menuItemIngredients := []MenuItemIngredient{}

	for i := 0; i < 10; i++ {
		ingredient := createRandomIngredient(t)
		arg := CreateMenuItemIngredientParams{
			FoodID:       menuItem.ID,
			IngredientID: ingredient.ID,
			Status:       "A",
			CreatedBy:    menuItem.CreatedBy,
		}

		res, err := testQueries.CreateMenuItemIngredient(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		menuItemIngredients = append(menuItemIngredients, res)
	}

	arg := GetAllActiveIngredientsFromFoodParams{
		FoodID: menuItem.ID,
		Limit:  5,
		Offset: 5,
	}

	res, err := testQueries.GetAllActiveIngredientsFromFood(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, res, 5)

	for _, item := range res {
		require.NotEmpty(t, item)
		require.Equal(t, item.FoodID, menuItem.ID)
		require.Equal(t, item.CreatedBy, menuItem.CreatedBy)
		require.Equal(t, item.Status, "A")
	}
}
