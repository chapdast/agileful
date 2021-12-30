package db

import (
	"context"
)

func (d database) UpdateViewCount(ctx context.Context, articleID int) error {
	result, err := d.conn.Exec(ctx, "update articles set view_count=view_count+1 WHERE id = $1", articleID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return nothingUpdated
	}
	return nil
}
