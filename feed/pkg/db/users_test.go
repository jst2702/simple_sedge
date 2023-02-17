package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"simplesedge.com/gokit/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:          util.RandomName(),
		HashedPassword: util.RandomString(20),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user)
	require.NotZero(t, user.CreatedAt)
	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	model1 := createRandomModel(t)
	model2, err := testQueries.GetModel(context.Background(), model1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, model2)
	require.Equal(t, model1, model2)
	require.WithinDuration(t, model1.CreatedAt, model2.CreatedAt, time.Second)
}

// func TestUpdateModel(t *testing.T) {
// 	model1 := createRandomModel(t)

// 	arg := UpdateModelParams{
// 		ID:          model1.ID,
// 		Name:        model1.Name,
// 		Description: model1.Description,
// 		Role:        util.RandomString(3),
// 	}

// 	model2, err := testQueries.UpdateModel(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, model2)
// 	require.Equal(t, model1.Name, model2.Name)
// 	require.Equal(t, model1.Description, model2.Description)
// 	require.NotEqual(t, model1.Role, model2.Role)
// }

// func TestDeleteModel(t *testing.T) {
// 	model1 := createRandomModel(t)
// 	err := testQueries.DeleteModel(context.Background(), model1.ID)
// 	require.NoError(t, err)

// 	model2, err := testQueries.GetModel(context.Background(), model1.ID)
// 	require.Error(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, model2)
// }

// func TestListModels(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomModel(t)
// 	}

// 	arg := ListModelsParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	models, err := testQueries.ListModels(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, models, 5)

// 	for _, model := range models {
// 		require.NotEmpty(t, model)
// 	}
// }
