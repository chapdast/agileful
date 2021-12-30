package db

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDatabase_Read(t *testing.T) {
	ctx := context.Background()
	d := getDBConn(ctx, t)
	clearAllArticles(ctx, t, d)

	testArticles := []*Article{}
	for i := 0; i <= 100; i++ {
		testArticles = append(testArticles, &Article{
			Writer: fmt.Sprintf("test_writer_%d", i),
			Body:   fmt.Sprintf("test body text of I:%d", i),
		})
		if err := d.Create(ctx, testArticles[i]); err != nil {
			t.Fatal(err)
		}
	}
	t.Run("get All no filter", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			arts, err := d.Read(ctx, OptionLimit(20))
			So(err, ShouldBeNil)
			So(len(arts), ShouldEqual, 20)
		})
	})
	t.Run("get exact no filter", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			arts, err := d.Read(ctx,
				OptionFilter("writer", Condition{
				Value:    "test_writer_45",
				Operator: OprEqual,
				}),
				OptionLimit(20))
			So(err, ShouldBeNil)
			So(len(arts), ShouldEqual, 1)
			So(arts[0].Writer , ShouldEqual,"test_writer_45")
		})
	})

	t.Run("get All writers with 2x in name", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			arts, err := d.Read(ctx,
				OptionFilter("writer", Condition{
					Value:    "writer_2",
					Operator: OprSubString,
				}),
				OptionLimit(20))
			So(err, ShouldBeNil)
			So(len(arts), ShouldEqual, 11)
		})
	})

	t.Run("get All writers with 2x in name DESC ORDER", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			arts, err := d.Read(ctx,
				OptionFilter("writer", Condition{
					Value:    "writer_2",
					Operator: OprSubString,
				}),
				OptionOrder("id", true),
				OptionLimit(20))
			So(err, ShouldBeNil)
			So(len(arts), ShouldEqual, 11)
			lastId := arts[0].Id
			for _, article := range arts[1:] {
				if article.Id > lastId {
					t.Fatal("order failed")
				}
				lastId = article.Id
			}
		})
	})
	t.Run("get All writers with 2x in name DESC ORDER PAGE 2", func(t *testing.T) {
		Convey(t.Name(), t, func() {
			arts, err := d.Read(ctx,
				OptionFilter("writer", Condition{
					Value:    "writer_2",
					Operator: OprSubString,
				}),
				OptionOrder("id", true),
				OptionLimit(8),)
			So(err, ShouldBeNil)
			So(len(arts), ShouldEqual, 8)
			arts, err = d.Read(ctx,
				OptionFilter("writer", Condition{
					Value:    "writer_2",
					Operator: OprSubString,
				}),
				OptionOffset(8),
				OptionOrder("id", true),
				OptionLimit(10),)
			So(err, ShouldBeNil)
			So(len(arts), ShouldEqual, 3)
		})
	})
}
