package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenameTx(t *testing.T) {
	model := createRandomModel(t)
	store := NewStore(testDB)
	result := store.RenameModelTx(context.Background(), RenameModelParams{model.ID})
	require.NoError(t, result.Err)
	model2, err := testQueries.GetModel(context.Background(), model.ID)
	require.NoError(t, err)
	require.Equal(t, model2.Name, "jaqapoe")
}
