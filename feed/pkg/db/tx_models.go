package db

import (
	"context"
	"database/sql"

	gokitSql "simplesedge.com/gokit/sql"
)

func (store *SQLStore) RenameModelTx(
	ctx context.Context,
	arg RenameModelParams) RenameModelResult {

	err := gokitSql.ExecTx(ctx, store.db, func(tx *sql.Tx) error {
		q := New(tx)
		_, err := q.UpdateModel(ctx, UpdateModelParams{
			ID: arg.ModelID, Name: "jaqapoe",
		})
		if err != nil {
			return err
		}
		return nil
	}, nil)

	return RenameModelResult{err}
}
