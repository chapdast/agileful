package db

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDatabase_Update(t *testing.T) {
	ctx := context.Background()
	d := getDBConn(ctx, t)

	t.Run("update success", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			clearAllArticles(ctx, t, d)
			id := int64(-1)
			So(d.Raw().QueryRow(ctx,
				`insert into articles(writer, body) 
VALUES ('test','test') RETURNING id`).Scan(&id), ShouldBeNil)
			update := &Article{
				Id:     id,
				Writer: "UPDATED",
				Body:   "UPDATED",
			}
			err := d.Update(ctx, update)
			So(err, ShouldBeNil)
			upArticle := &Article{}
			So(d.Raw().QueryRow(ctx,
				`SELECT id, writer, body FROM articles WHERE id=$1`, id).
				Scan(&upArticle.Id, &upArticle.Writer, &upArticle.Body), ShouldBeNil)
			So(upArticle, ShouldResemble, update)
		})
	})
	t.Run("update none existed", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			clearAllArticles(ctx, t, d)
			err := d.Update(ctx, &Article{Id: 12})
			So(err, ShouldBeError, nothingUpdated)
		})
	})

	t.Run("update error", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			clearAllArticles(ctx, t, d)
			ctx, cf := context.WithCancel(ctx)
			cf()
			err := d.Update(ctx, &Article{
				Id:     10,
				Writer: "test",
				Body:   "test",
			})
			So(err, ShouldBeError, context.Canceled)
		})
	})

}
