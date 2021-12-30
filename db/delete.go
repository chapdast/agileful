package db

import (
	"context"
	"errors"
)

var nothingDeleted = errors.New("no article has deleted")

func (d database) Delete(ctx context.Context, articleID int) error {
	result, err := d.conn.Exec(ctx, "DELETE FROM articles where id=$1", articleID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return nothingDeleted
	}
	return nil
}
