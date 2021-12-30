package db

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDatabase_Delete(t *testing.T) {
	ctx := context.Background()
	d := getDBConn(ctx, t)
	t.Run("delete existed", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			clearAllArticles(ctx, t, d)
			id := -1
			So(d.Raw().QueryRow(ctx,
				`insert into articles(writer, body) 
VALUES ('test','test') RETURNING id`).Scan(&id), ShouldBeNil)
			So(id, ShouldNotEqual, -1)
			err := d.Delete(ctx, id)
			So(err, ShouldBeNil)
		})
	})
	t.Run("delete none existed", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			clearAllArticles(ctx, t, d)
			id := -1
			err := d.Delete(ctx, id)
			So(err, ShouldBeError, nothingDeleted)
		})
	})

}
