package db

import (
	"context"
	"database/sql"
)

type RenameModelParams struct {
	ModelID int64 `json:"model_id"`
}

type RenameModelResult struct {
	Err error `json:"error"`
}

// Using a store interface allows for easier mocking.
type Store interface {
	Querier
	RenameModelTx(context.Context, RenameModelParams) RenameModelResult
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
