package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vincentyeungg/golang-menu-creator/util"
)

func createRandomIngredient(t *testing.T) Ingredient {
	owner := util.RandomName()

	arg := CreateIngredientParams{
		Name:        util.RandomName(),
		BrandName:   util.RandomName(),
		Description: util.RandomDescription(),
		Status:      "A",
		CreatedBy:   owner,
		UpdatedBy:   owner,
	}

	ingredient, err := testQueries.CreateIngredient(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, ingredient)

	require.Equal(t, arg.Name, ingredient.Name)
	require.Equal(t, arg.BrandName, ingredient.BrandName)
	require.Equal(t, arg.Description, ingredient.Description)
	require.Equal(t, arg.Status, ingredient.Status)
	require.Equal(t, arg.CreatedBy, ingredient.CreatedBy)
	require.Equal(t, arg.UpdatedBy, ingredient.UpdatedBy)

	require.NotZero(t, ingredient.ID)
	require.NotZero(t, ingredient.CreatedAt)

	return ingredient
}

func TestCreateIngredient(t *testing.T) {
	createRandomIngredient(t)
}

func TestGetIngredient(t *testing.T) {
	ingredient1 := createRandomIngredient(t)

	arg := GetIngredientParams{
		ID:     ingredient1.ID,
		Status: ingredient1.Status,
	}

	res, err := testQueries.GetIngredient(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, ingredient1.ID, res.ID)
}

func TestUpdateIngredient(t *testing.T) {
	ingredient1 := createRandomIngredient(t)

	arg := UpdateIngredientParams{
		Name:        util.RandomName(),
		BrandName:   util.RandomName(),
		Description: util.RandomDescription(),
		ID:          ingredient1.ID,
	}

	res, err := testQueries.UpdateIngredient(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, ingredient1.ID, res.ID)
	require.Equal(t, arg.Name, res.Name)
	require.Equal(t, arg.BrandName, res.BrandName)
	require.Equal(t, arg.Description, res.Description)
	require.Equal(t, ingredient1.Status, res.Status)
	require.Equal(t, ingredient1.CreatedBy, res.CreatedBy)
	require.Equal(t, ingredient1.UpdatedBy, res.UpdatedBy)
}

func TestDeleteIngredient(t *testing.T) {
	ingredient1 := createRandomIngredient(t)
	err := testQueries.DeleteIngredient(context.Background(), ingredient1.ID)
	require.NoError(t, err)

	// arg := GetIngredientParams{
	// 	ID:     ingredient1.ID,
	// 	Status: ingredient1.Status,
	// }

	// res, err := testQueries.GetIngredient(context.Background(), arg)

	// require.NoError(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	// require.Empty(t, res)
}

func TestListIngredients(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomIngredient(t)
	}

	arg := GetAllIngredientParams{
		Limit:  5,
		Offset: 5,
	}

	ingredients, err := testQueries.GetAllIngredient(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, ingredients, 5)

	for _, ingredient := range ingredients {
		require.NotEmpty(t, ingredient)
	}
}

func TestListActiveIngredients(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomIngredient(t)
	}

	arg := GetAllActiveIngredientsParams{
		Limit:  5,
		Offset: 5,
	}

	ingredients, err := testQueries.GetAllActiveIngredients(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, ingredients, 5)

	for _, ingredient := range ingredients {
		require.NotEmpty(t, ingredient)
		require.Equal(t, ingredient.Status, "A")
	}
}
