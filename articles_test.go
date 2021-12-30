package agileful

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chapdast/agileful/db"
	"github.com/chapdast/agileful/mock_db"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"math/rand"
	"net/http"
	"testing"
)

func TestArticles(t *testing.T) {

	port := rand.Intn(3000) + 3000

	testCase := map[string]struct {
		req    string
		resp   []db.Article
		code   int
		mocker func(conn *mock_db.MockDatabase)
	}{
		"no filter only default page size": {
			req: fmt.Sprintf(srvURL, port),
			resp: []db.Article{
				{
					Id:         1,
					Writer:     "test",
					Body:       "test",
					CreateDate: 0,
					ViewCount:  10,
				},
				{
					Id:         2,
					Writer:     "test_2",
					Body:       "test_3",
					CreateDate: 0,
					ViewCount:  13,
				},
			},
			code: 200,
			mocker: func(conn *mock_db.MockDatabase) {
				conn.EXPECT().Read(gomock.Any(), gomock.Any()).
					Return([]*db.Article{
						{
							Id:         1,
							Writer:     "test",
							Body:       "test",
							CreateDate: 0,
							ViewCount:  10,
						},
						{
							Id:         2,
							Writer:     "test_2",
							Body:       "test_3",
							CreateDate: 0,
							ViewCount:  13,
						},
					}, nil)
			},
		},
		"filter writers only default page size": {
			req: srvURL + "?writer=test_2",
			resp: []db.Article{
				{
					Id:         2,
					Writer:     "test_2",
					Body:       "test_3",
					CreateDate: 0,
					ViewCount:  13,
				},
			},
			code: 200,
			mocker: func(conn *mock_db.MockDatabase) {
				conn.EXPECT().Read(gomock.Any(), gomock.Any(), gomock.Any()).
					Return([]*db.Article{
						{
							Id:         2,
							Writer:     "test_2",
							Body:       "test_3",
							CreateDate: 0,
							ViewCount:  13,
						},
					}, nil)
			},
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	conn := mock_db.NewMockDatabase(ctrl)
	app, err := Run(conn)
	if err != nil {
		t.Fatal(err)
	}
	go runServer(ctx, t, app, port)
	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			if test.mocker != nil {
				test.mocker(conn)
			}
			Convey(t.Name(), t, func() {
				resp, err := http.Get(fmt.Sprintf(srvURL, port))
				if err != nil {
					t.Fatal("ERROR:", err)
				}
				var result []db.Article
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Fatal(err)
				}
				defer resp.Body.Close()
				if err := json.Unmarshal(body, &result); err != nil {
					t.Log(body)
					t.Fatal(err)
				}
				So(result, ShouldResemble, test.resp)
			})

		})
	}

}
