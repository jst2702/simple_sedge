package db

import (
	"context"
	"database/sql"

	gokitSql "simplesedge.com/gokit/sql"
)

func (store Store) RenameModelTx(
	ctx context.Context,
	arg RenameModelParams) RenameModelResult {

	err := gokitSql.ExecTx(ctx, store.db, func(tx *sql.Tx) error {
		var err error
		_, err = store.UpdateModel(ctx, UpdateModelParams{
			ID: arg.ModelID, Name: "jaqapoe",
		})
		if err != nil {
			return err
		}
		return nil
	}, nil)

	return RenameModelResult{err}
}
