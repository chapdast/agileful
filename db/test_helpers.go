package db

import (
	"context"
	"testing"
)

func getDBConn(ctx context.Context, t *testing.T) Database {
	d, err := NewDB(ctx, GetDBLinkFromEnv())
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func clearAllArticles(ctx context.Context, t *testing.T, d Database) {
	_, err := d.Raw().Exec(ctx, "DELETE FROM articles")
	if err != nil {
		t.Fatal(err)
	}
}
