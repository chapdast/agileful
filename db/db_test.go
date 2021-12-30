package db

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestNewDB(t *testing.T) {
	t.Run("New Conn no error", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			ctx := context.Background()
			db, err := NewDB(ctx, GetDBLinkFromEnv())
			So(err, ShouldBeNil)
			So(db, ShouldNotBeNil)
			So(db.Raw(), ShouldNotBeNil)
		})
	})
	t.Run("error  context", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			ctx, cf := context.WithCancel(context.Background())
			cf()
			db, err := NewDB(ctx, GetDBLinkFromEnv())
			So(err, ShouldBeError, context.Canceled)
			So(db, ShouldBeNil)
		})
	})
	t.Run("error on New connection", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			ctx := context.Background()
			db, err := NewDB(ctx, "INVALID_LINK")
			So(strings.HasSuffix(err.Error(), "(invalid dsn)"), ShouldBeTrue)
			So(db, ShouldBeNil)
		})
	})

}
