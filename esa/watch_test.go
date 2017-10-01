package esa

import (
	"context"
	"net/http"
	"testing"
)

func TestClientImpl_ListPostWatchers(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/posts/2312/watchers",
		query:  "page=2&per_page=30",
		res: `
		{
			"watchers": [
			{
				"created_at": "2016-05-05T11:40:53+09:00",
				"user": {
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				}
			}
			],
			"prev_page": null,
			"next_page": null,
			"total_count": 1,
			"page": 1,
			"per_page": 20,
			"max_per_page": 100
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.ListPostWatchers(
				context.Background(),
				2312,
				2,
				30,
			)
		},
	})
}

func TestClientImpl_WatchPost(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/posts/2312/watch",
		do: func(client Client) (interface{}, error) {
			return nil, client.WatchPost(
				context.Background(),
				2312,
			)
		},
	})
}

func TestClientImpl_UnwatchPost(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodDelete,
		spath:  "/v1/teams/docs/posts/2312/watch",
		do: func(client Client) (interface{}, error) {
			return nil, client.UnwatchPost(
				context.Background(),
				2312,
			)
		},
	})
}
