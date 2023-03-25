package db

import (
	"context"
	"database/sql"

	gokitSql "simplesedge.com/gokit/sql"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := gokitSql.ExecTx(ctx, store.db, func(tx *sql.Tx) error {
		var err error
		q := New(tx) // queries created from the tx obj can be rolled back.
		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.User)
	}, nil)

	return result, err
}
