package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"simplesedge.com/gokit/util"
)

func createRandomDocument(t *testing.T) Document {
	arg := CreateDocumentParams{
		Guid:        util.RandomString(45),
		Url:         util.RandomString(30),
		Site:        util.RandomString(30),
		SiteFull:    util.RandomString(30),
		SiteSection: util.RandomString(30),
		Headline:    util.RandomString(30),
		Title:       util.RandomString(30),
		Body:        util.RandomString(30),
		Ticker:      util.RandomNullString(4),
		PublishedAt: util.RandomTimestamp(),
		Tickers: []string{
			util.RandomString(4),
			util.RandomString(4),
		},
	}

	doc, err := testQueries.CreateDocument(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, doc)
	require.Equal(t, arg.Headline, doc.Headline)
	require.Equal(t, arg.Headline, doc.Headline)

	require.NotEmpty(t, doc.Guid)
	require.NotZero(t, doc.CreatedAt)
	return doc
}
func TestCreateDocument(t *testing.T) {
	createRandomDocument(t)
}

func TestGetDocument(t *testing.T) {
	doc1 := createRandomDocument(t)
	doc2, err := testQueries.GetDocument(context.Background(), doc1.Guid)
	require.NoError(t, err)
	require.NotEmpty(t, doc2)
	require.Equal(t, doc1, doc2)
	require.WithinDuration(t, doc1.CreatedAt, doc2.CreatedAt, time.Second)
}

func TestUpdateDocument(t *testing.T) {
	doc1 := createRandomDocument(t)

	arg := UpdateDocumentParams{
		Guid:   doc1.Guid,
		Ticker: util.RandomNullString(4),
		Tickers: []string{
			util.RandomString(4),
			util.RandomString(4),
		},
	}

	doc2, err := testQueries.UpdateDocument(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, doc2)
	require.Equal(t, doc2.Headline, doc2.Headline)
	require.Equal(t, doc2.Title, doc2.Title)
	require.NotEqual(t, doc1.Ticker, doc2.Ticker)
	require.NotEqual(t, doc1.Tickers, doc2.Tickers)
}

func TestDeleteDocument(t *testing.T) {
	doc1 := createRandomDocument(t)
	err := testQueries.DeleteDocuemnt(context.Background(), doc1.Guid)
	require.NoError(t, err)

	doc2, err := testQueries.GetDocument(context.Background(), doc1.Guid)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, doc2)
}

func TestListDocuments(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomDocument(t)
	}

	arg := ListDocumentsParams{
		Limit:  5,
		Offset: 5,
	}

	models, err := testQueries.ListDocuments(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, models, 5)

	for _, model := range models {
		require.NotEmpty(t, model)
	}
}
