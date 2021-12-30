package db

import "context"

func (d database) Create(ctx context.Context, article *Article) error {
	_, err := d.conn.Exec(ctx,
		"insert into articles(writer, body) VALUES($1, $2)", article.Writer, article.Body)
	return err
}
