package db

import (
	"database/sql"
)

type RenameModelParams struct {
	ModelID int64 `json:"model_id"`
}

type RenameModelResult struct {
	Err error `json:"error"`
}

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
