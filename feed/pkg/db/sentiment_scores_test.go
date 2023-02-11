package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"simplesedge.com/gokit/util"
)

func createRandomScore(t *testing.T) SentimentScore {
	model := createRandomModel(t)
	doc := createRandomDocument(t)
	arg := CreateScoreParams{
		ModelID:      model.ID,
		DocumentGuid: doc.Guid,
		Sentiment:    util.RandomInt(0, 100),
		Confidence:   util.RandomInt(0, 100),
	}

	score, err := testQueries.CreateScore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, score)
	require.Equal(t, arg.Sentiment, score.Sentiment)
	require.Equal(t, arg.Confidence, score.Confidence)
	require.Equal(t, arg.DocumentGuid, doc.Guid)
	require.Equal(t, arg.ModelID, model.ID)

	require.NotZero(t, score.ID)
	require.NotZero(t, score.CreatedAt)
	return score
}

func TestCreateScore(t *testing.T) {
	createRandomScore(t)
}

func TestGetScore(t *testing.T) {
	Score1 := createRandomScore(t)
	Score2, err := testQueries.GetScore(context.Background(), Score1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, Score2)
	require.Equal(t, Score1, Score2)
	require.WithinDuration(t, Score1.CreatedAt, Score2.CreatedAt, time.Second)
}

func TestUpdateScore(t *testing.T) {
	Score1 := createRandomScore(t)

	arg := UpdateScoreParams{
		ID:         Score1.ID,
		Sentiment:  util.RandomInt(0, 100),
		Confidence: util.RandomInt(0, 100),
	}

	Score2, err := testQueries.UpdateScore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Score2)
	require.Equal(t, Score1.ModelID, Score2.ModelID)
	require.Equal(t, Score1.DocumentGuid, Score2.DocumentGuid)
	require.NotEqual(t, Score1.Sentiment, Score2.Sentiment)
	require.NotEqual(t, Score1.Confidence, Score2.Confidence)
}

func TestDeleteScore(t *testing.T) {
	Score1 := createRandomScore(t)
	err := testQueries.DeleteScore(context.Background(), Score1.ID)
	require.NoError(t, err)

	Score2, err := testQueries.GetScore(context.Background(), Score1.ID)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, Score2)
}

func TestListScores(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomScore(t)
	}

	arg := ListScoresParams{
		Limit:  5,
		Offset: 5,
	}

	Scores, err := testQueries.ListScores(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, Scores, 5)

	for _, Score := range Scores {
		require.NotEmpty(t, Score)
	}
}
