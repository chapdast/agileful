package db

import (
	"context"
	"errors"
)

var nothingUpdated = errors.New("no article has updated")

func (d database) Update(ctx context.Context, article *Article) error {
	result, err := d.conn.Exec(ctx,
		`UPDATE articles SET writer=$2, body=$3 WHERE id = $1;`, article.Id,
		article.Writer, article.Body)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return nothingUpdated
	}
	return nil
}
