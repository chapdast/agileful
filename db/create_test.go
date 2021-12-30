package db

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDatabase_Create(t *testing.T) {
	ctx := context.Background()
	d := getDBConn(ctx, t)
	Convey(t.Name(), t, func() {
		clearAllArticles(ctx, t, d)
		count := -1
		So(d.Raw().
			QueryRow(ctx, "SELECT count(id) FROM articles").
			Scan(&count), ShouldBeNil)
		So(count, ShouldEqual, 0)

		err := d.Create(ctx, &Article{
			Writer: "TestWriter",
			Body:   "lorem ipsum body",
		})
		So(err, ShouldBeNil)
	})

}
