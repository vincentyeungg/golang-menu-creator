package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vincentyeungg/golang-menu-creator/util"
)

// func to create a random menu item for testing
func createRandomMenuItem(t *testing.T) MenuItem {
	owner := util.RandomName()

	arg := CreateMenuItemParams{
		Name:        util.RandomName(),
		Price:       util.RandomPrice(),
		Description: util.RandomDescription(),
		Status:      "A",
		CreatedBy:   owner,
		UpdatedBy:   owner,
	}

	menuItem, err := testQueries.CreateMenuItem(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, menuItem)

	require.Equal(t, arg.Name, menuItem.Name)
	require.Equal(t, arg.Price, menuItem.Price)
	require.Equal(t, arg.Description, menuItem.Description)
	require.Equal(t, arg.Status, menuItem.Status)
	require.Equal(t, arg.CreatedBy, menuItem.CreatedBy)
	require.Equal(t, arg.UpdatedBy, menuItem.UpdatedBy)

	require.NotZero(t, menuItem.ID)
	require.NotZero(t, menuItem.CreatedAt)

	return menuItem
}

func TestCreateMenuItem(t *testing.T) {
	createRandomMenuItem(t)
}

func TestGetMenuItem(t *testing.T) {
	menuItem1 := createRandomMenuItem(t)
	arg := GetMenuItemParams{
		ID:     menuItem1.ID,
		Status: menuItem1.Status,
	}
	res, err := testQueries.GetMenuItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, menuItem1.ID, res.ID)
	require.Equal(t, menuItem1.Name, res.Name)
	require.Equal(t, menuItem1.Description, res.Description)
	require.Equal(t, menuItem1.Price, res.Price)
	require.Equal(t, menuItem1.Status, res.Status)
	require.Equal(t, menuItem1.CreatedBy, res.CreatedBy)
	require.Equal(t, menuItem1.UpdatedBy, res.UpdatedBy)
}

func TestUpdateMenuItem(t *testing.T) {
	menuItem1 := createRandomMenuItem(t)

	arg := UpdateMenuItemParams{
		ID: menuItem1.ID,
		Name: util.RandomName(),
		Description: util.RandomDescription(),
		Price: util.RandomPrice(),
	}

	res, err := testQueries.UpdateMenuItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, menuItem1.ID, res.ID)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Description, res.Description)
	require.Equal(t, arg.Price, res.Price)
	require.Equal(t, menuItem1.Status, res.Status)
	require.Equal(t, menuItem1.CreatedBy, res.CreatedBy)
	require.Equal(t, menuItem1.UpdatedBy, res.UpdatedBy)
}

func TestDeleteMenuItem(t *testing.T) {
	menuItem1 := createRandomMenuItem(t)
	err := testQueries.DeleteMenuItem(context.Background(), menuItem1.ID)
	require.NoError(t, err)

	arg := GetMenuItemParams{
		ID:     menuItem1.ID,
		Status: menuItem1.Status,
	}
	res, err := testQueries.GetMenuItem(context.Background(), arg)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}

func TestListMenuItems(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMenuItem(t)
	}

	arg := GetAllMenuItemsParams{
		Limit: 5,
		Offset: 5,
	}

	menuItems, err := testQueries.GetAllMenuItems(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, menuItems, 5)

	for _, menuItem := range menuItems {
		require.NotEmpty(t, menuItem)
	}
}

func TestListActiveMenuItems(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMenuItem(t)
	}

	arg := GetAllActiveItemsParams{
		Limit: 5,
		Offset: 5,
	}

	menuItems, err := testQueries.GetAllActiveItems(context.Background(), arg)
	require.NoError(t, err)

	for _, menuItem := range menuItems {
		require.NotEmpty(t, menuItem)
		require.Equal(t, menuItem.Status, "A")
	}
}