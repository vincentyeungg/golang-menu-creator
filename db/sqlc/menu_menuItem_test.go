package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomMenuMenuItem(t *testing.T) MenuMenuItem {
	// create menu and menu item
	menu := createRandomMenu(t)
	menuItem := createRandomMenuItem(t)

	arg := CreateMenuMenuItemParams{
		MenuID:    menu.ID,
		FoodID:    menuItem.ID,
		Status:    "A",
		CreatedBy: menu.CreatedBy,
		UpdatedBy: menu.CreatedBy,
	}

	res, err := testQueries.CreateMenuMenuItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, arg.MenuID, res.MenuID)
	require.Equal(t, arg.FoodID, res.FoodID)
	require.Equal(t, arg.Status, res.Status)
	require.Equal(t, arg.CreatedBy, res.CreatedBy)
	require.Equal(t, arg.UpdatedBy, res.UpdatedBy)

	require.NotZero(t, res.ID)
	require.NotZero(t, res.CreatedAt)

	return res
}

func TestCreateRandomMenuMenuItem(t *testing.T) {
	CreateRandomMenuMenuItem(t)
}

func TestGetMenuMenuItem(t *testing.T) {
	menuMenuItem := CreateRandomMenuMenuItem(t)

	res, err := testQueries.GetMenuMenuItem(context.Background(), menuMenuItem.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, menuMenuItem.ID, res.ID)
	require.Equal(t, menuMenuItem.FoodID, res.FoodID)
	require.Equal(t, menuMenuItem.FoodID, res.FoodID)
	require.Equal(t, menuMenuItem.CreatedBy, res.CreatedBy)
	require.Equal(t, menuMenuItem.UpdatedBy, res.UpdatedBy)
	require.Equal(t, menuMenuItem.Status, res.Status)
}

func TestGetActiveItemFromMenu(t *testing.T) {
	menuMenuItem := CreateRandomMenuMenuItem(t)

	res, err := testQueries.GetMenuMenuItem(context.Background(), menuMenuItem.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, menuMenuItem.ID, res.ID)
	require.Equal(t, menuMenuItem.FoodID, res.FoodID)
	require.Equal(t, menuMenuItem.FoodID, res.FoodID)
	require.Equal(t, menuMenuItem.CreatedBy, res.CreatedBy)
	require.Equal(t, menuMenuItem.UpdatedBy, res.UpdatedBy)
	require.Equal(t, res.Status, "A")
}

func TestDeleteMenuMenuItem(t *testing.T) {
	menuMenuItem := CreateRandomMenuMenuItem(t)

	arg := DeleteMenuItemFromMenuParams{
		MenuID:    menuMenuItem.MenuID,
		FoodID:    menuMenuItem.FoodID,
		CreatedBy: menuMenuItem.CreatedBy,
	}

	err := testQueries.DeleteMenuItemFromMenu(context.Background(), arg)
	res, err := testQueries.GetMenuMenuItem(context.Background(), menuMenuItem.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}

func TestGetAllItemsFromMenu(t *testing.T) {
	menu := createRandomMenu(t)
	menuMenuItems := []MenuMenuItem{}

	for i := 0; i < 10; i++ {
		menuItem := createRandomMenuItem(t)
		arg := CreateMenuMenuItemParams{
			FoodID:    menuItem.ID,
			MenuID:    menu.ID,
			Status:    "A",
			CreatedBy: menu.CreatedBy,
			UpdatedBy: menu.UpdatedBy,
		}

		res, err := testQueries.CreateMenuMenuItem(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		menuMenuItems = append(menuMenuItems, res)
	}

	arg := GetAllItemsFromMenuParams{
		MenuID: menu.ID,
		Limit:  5,
		Offset: 5,
	}

	res, err := testQueries.GetAllItemsFromMenu(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, res, 5)

	for _, item := range res {
		require.NotEmpty(t, item)
		require.Equal(t, item.MenuID, menu.ID)
	}
}

func TestGetAllActiveItemsFromMenu(t *testing.T) {
	menu := createRandomMenu(t)
	menuMenuItems := []MenuMenuItem{}

	for i := 0; i < 10; i++ {
		menuItem := createRandomMenuItem(t)
		arg := CreateMenuMenuItemParams{
			FoodID:    menuItem.ID,
			MenuID:    menu.ID,
			Status:    "A",
			CreatedBy: menu.CreatedBy,
			UpdatedBy: menu.UpdatedBy,
		}

		res, err := testQueries.CreateMenuMenuItem(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		menuMenuItems = append(menuMenuItems, res)
	}

	arg := GetAllActiveItemsFromMenuParams{
		MenuID: menu.ID,
		Limit:  5,
		Offset: 5,
	}

	res, err := testQueries.GetAllActiveItemsFromMenu(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, res, 5)

	for _, item := range res {
		require.NotEmpty(t, item)
		require.Equal(t, item.MenuID, menu.ID)
		require.Equal(t, item.Status, "A")
	}
}
