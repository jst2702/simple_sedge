package sql

import (
	"context"
	"database/sql"
	"fmt"
)

func ExecTx(
	ctx context.Context,
	db *sql.DB,
	fn func(*sql.Tx) error,
	opts *sql.TxOptions) error {

	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
