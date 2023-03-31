package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"simplesedge.com/gokit/util"
)

func createRandomApiKey(t *testing.T) ApiKey {
	keyString, err := util.NewKeyString()
	require.NoError(t, err)

	key, err := testQueries.CreateApiKey(context.Background(), keyString)
	require.NoError(t, err)
	require.NotEmpty(t, key)
	require.True(t, key.Active)
	return key
}
func TestCreateApiKey(t *testing.T) {
	createRandomApiKey(t)
}

func TestGetApiKey(t *testing.T) {
	key1 := createRandomApiKey(t)
	key2, err := testQueries.GetApiKey(context.Background(), key1.ApiKey)
	require.NoError(t, err)
	require.NotEmpty(t, key2)
	require.Equal(t, key1, key2)
	require.WithinDuration(t, key1.CreatedAt, key2.CreatedAt, time.Second)
}

func TestUpdateApiKey(t *testing.T) {
	key1 := createRandomApiKey(t)

	key2, err := testQueries.DisableApiKey(context.Background(), key1.ApiKey)
	require.NoError(t, err)

	require.Equal(t, key1.ApiKey, key2.ApiKey)
	require.False(t, key2.Active)
}
