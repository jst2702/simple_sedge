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
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomString(20),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
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
