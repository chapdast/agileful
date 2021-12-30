package db

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDatabase_UpdateViewCount(t *testing.T) {
	ctx := context.Background()
	d := getDBConn(ctx, t)
	t.Run("existed", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			clearAllArticles(ctx, t, d)
			id := -1
			So(d.Raw().QueryRow(ctx,
				`insert into articles(writer, body) 
VALUES ('test','test') RETURNING id`).Scan(&id), ShouldBeNil)

			err := d.UpdateViewCount(ctx, id)
			So(err, ShouldBeNil)

			err = d.UpdateViewCount(ctx, id)
			So(err, ShouldBeNil)

			count := -1
			So(d.Raw().QueryRow(ctx, "SELECT view_count FROM articles WHERE id=$1", id).
				Scan(&count), ShouldBeNil)
			So(count, ShouldEqual, 2)
		})
	})
	t.Run("none existed", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			clearAllArticles(ctx, t, d)
			id := -1
			err := d.UpdateViewCount(ctx, id)
			So(err, ShouldBeError, nothingUpdated)
		})
	})

}
