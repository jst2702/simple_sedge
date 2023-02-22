// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	CreateDocument(ctx context.Context, arg CreateDocumentParams) (Document, error)
	CreateModel(ctx context.Context, arg CreateModelParams) (Model, error)
	CreateScore(ctx context.Context, arg CreateScoreParams) (SentimentScore, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteDocuemnt(ctx context.Context, guid string) error
	DeleteModel(ctx context.Context, id int64) error
	DeleteScore(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, email string) error
	GetDocument(ctx context.Context, guid string) (Document, error)
	GetModel(ctx context.Context, id int64) (Model, error)
	GetScore(ctx context.Context, id int64) (SentimentScore, error)
	GetUser(ctx context.Context, email string) (User, error)
	ListDocuments(ctx context.Context, arg ListDocumentsParams) ([]Document, error)
	ListModels(ctx context.Context, arg ListModelsParams) ([]Model, error)
	ListScores(ctx context.Context, arg ListScoresParams) ([]SentimentScore, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateDocument(ctx context.Context, arg UpdateDocumentParams) (Document, error)
	UpdateModel(ctx context.Context, arg UpdateModelParams) (Model, error)
	UpdateScore(ctx context.Context, arg UpdateScoreParams) (SentimentScore, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
