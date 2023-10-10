package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vincentyeungg/golang-menu-creator/util"
)

func createRandomMenu(t *testing.T) Menu {
	owner := util.RandomName()

	arg := CreateMenuParams{
		Name: util.RandomName(),
		Description: util.RandomDescription(),
		Status: "A",
		CreatedBy: owner,
		UpdatedBy: owner,
	}

	menu, err := testQueries.CreateMenu(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, menu)

	require.Equal(t, arg.Name, menu.Name)
	require.Equal(t, arg.Description, menu.Description)
	require.Equal(t, arg.Status, menu.Status)
	require.Equal(t, arg.CreatedBy, menu.CreatedBy)
	require.Equal(t, arg.UpdatedBy, menu.UpdatedBy)

	return menu
}

func TestCreateMenu(t *testing.T) {
	createRandomMenu(t)
}

func TestGetMenu(t *testing.T) {
	menu1 := createRandomMenu(t)
	res, err := testQueries.GetMenu(context.Background(), menu1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, menu1.ID, res.ID)
	require.Equal(t, menu1.Name, res.Name)
	require.Equal(t, menu1.Description, res.Description)
	require.Equal(t, menu1.Status, res.Status)
	require.Equal(t, menu1.CreatedBy, res.CreatedBy)
	require.Equal(t, menu1.UpdatedBy, res.UpdatedBy)
}

func TestUpdateMenu(t *testing.T) {
	menu1 := createRandomMenu(t)

	arg := UpdateMenuParams{
		ID: menu1.ID,
		Name: menu1.Name,
		Description: menu1.Description,
	}

	res, err := testQueries.UpdateMenu(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, menu1.ID, res.ID)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.Description, res.Description)
	require.Equal(t, menu1.Status, res.Status)
	require.Equal(t, menu1.CreatedBy, res.CreatedBy)
	require.Equal(t, menu1.UpdatedBy, res.UpdatedBy)
}

func TestDeleteMenu(t *testing.T) {
	menu1 := createRandomMenu(t)
	err := testQueries.DeleteMenu(context.Background(), menu1.ID)
	require.NoError(t, err)

	res, err := testQueries.GetMenu(context.Background(), menu1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}

func TestListMenus(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMenu(t)
	}

	arg := GetAllMenusParams{
		Limit: 5,
		Offset: 5,
	}

	menus, err := testQueries.GetAllMenus(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, menus, 5)

	for _, menu := range menus {
		require.NotEmpty(t, menu)
	}
}

func TestListAllActiveMenus(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMenu(t)
	}

	arg := GetAllActiveMenusParams{
		Limit: 5,
		Offset: 5,
	}

	menus, err := testQueries.GetAllActiveMenus(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, menus, 5)

	for _, menu := range menus {
		require.NotEmpty(t, menu)
		require.Equal(t, menu.Status, "A")
	}
}
