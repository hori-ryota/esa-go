package esa

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestClientImpl_ListPostStargazers(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/posts/2312/stargazers",
		query:  "page=2&per_page=30",
		res: `
		{
			"stargazers": [
			{
				"created_at": "2016-05-05T11:40:54+09:00",
				"body": null,
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
			return client.ListPostStargazers(
				context.Background(),
				2312,
				2,
				30,
			)
		},
	})
}

func TestClientImpl_StarPost(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/posts/2312/star",
		body: `
		{"body":"foo bar"}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			p := StarParam{}
			if err := json.Unmarshal([]byte(body), &p); err != nil {
				return nil, err
			}
			return nil, client.StarPost(
				context.Background(),
				2312,
				p,
			)
		},
	})
}

func TestClientImpl_UnstarPost(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodDelete,
		spath:  "/v1/teams/docs/posts/2312/star",
		do: func(client Client) (interface{}, error) {
			return nil, client.UnstarPost(
				context.Background(),
				2312,
			)
		},
	})
}

func TestClientImpl_ListCommentStargazers(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/comments/123/stargazers",
		query:  "page=2&per_page=30",
		res: `
		{
			"stargazers": [
			{
				"created_at": "2016-05-05T11:40:54+09:00",
				"body": null,
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
			return client.ListCommentStargazers(
				context.Background(),
				123,
				2,
				30,
			)
		},
	})
}

func TestClientImpl_StarComment(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/comments/123/star",
		body: `
		{"body":"foo bar"}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			p := StarParam{}
			if err := json.Unmarshal([]byte(body), &p); err != nil {
				return nil, err
			}
			return nil, client.StarComment(
				context.Background(),
				123,
				p,
			)
		},
	})
}

func TestClientImpl_UnstarComment(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodDelete,
		spath:  "/v1/teams/docs/comments/123/star",
		do: func(client Client) (interface{}, error) {
			return nil, client.UnstarComment(
				context.Background(),
				123,
			)
		},
	})
}
