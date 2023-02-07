package data

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"simplesedge.com/gokit/util"
)

func createRandomModel(t *testing.T) Model {
	arg := CreateModelParams{
		Name:        util.RandomName(),
		Role:        util.RandomString(3),
		Description: util.RandomNullString(30),
	}

	model, err := testQueries.CreateModel(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, model)
	require.Equal(t, arg.Name, model.Name)
	require.Equal(t, arg.Role, model.Role)

	require.NotZero(t, model.ID)
	require.NotZero(t, model.CreatedAt)
	return model
}
func TestCreateModel(t *testing.T) {
	createRandomModel(t)
}

func TestGetModel(t *testing.T) {
	model1 := createRandomModel(t)
	model2, err := testQueries.GetModel(context.Background(), model1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, model2)
	require.Equal(t, model1, model2)
	require.WithinDuration(t, model1.CreatedAt, model2.CreatedAt, time.Second)
}

func TestUpdateModel(t *testing.T) {
	model1 := createRandomModel(t)

	arg := UpdateModelParams{
		ID:          model1.ID,
		Name:        model1.Name,
		Description: model1.Description,
		Role:        util.RandomString(3),
	}

	model2, err := testQueries.UpdateModel(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, model2)
	require.Equal(t, model1.Name, model2.Name)
	require.Equal(t, model1.Description, model2.Description)
	require.NotEqual(t, model1.Role, model2.Role)
}

func TestDeleteModel(t *testing.T) {
	model1 := createRandomModel(t)
	err := testQueries.DeleteModel(context.Background(), model1.ID)
	require.NoError(t, err)

	model2, err := testQueries.GetModel(context.Background(), model1.ID)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, model2)
}

func TestListModels(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomModel(t)
	}

	arg := ListModelsParams{
		Limit:  5,
		Offset: 5,
	}

	models, err := testQueries.ListModels(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, models, 5)

	for _, model := range models {
		require.NotEmpty(t, model)
	}
}
