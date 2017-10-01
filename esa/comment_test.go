package esa

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestClientImpl_ListComments(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/posts/2/comments",
		query:  "page=2&per_page=30",
		res: `
		{
			"comments": [
			{
				"id": 1,
				"body_md": "(大事)",
				"body_html": "<p>(大事)</p>",
				"created_at": "2014-05-10T12:45:42+09:00",
				"updated_at": "2014-05-18T23:02:29+09:00",
				"url": "https://docs.esa.io/posts/2#comment-1",
				"created_by": {
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"stargazers_count": 0,
				"star": false
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
			return client.ListComments(
				context.Background(),
				2,
				2,
				30,
			)
		},
	})
}

func TestClientImpl_GetComment(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/comments/13",
		res: `
		{
			"id": 13,
			"body_md": "読みたい",
			"body_html": "<p>読みたい</p>",
			"created_at": "2014-05-13T16:17:42+09:00",
			"updated_at": "2014-05-18T23:02:29+09:00",
			"url": "https://docs.esa.io/posts/13#comment-13",
			"created_by": {
				"name": "TAEKO AKATSUKA",
				"screen_name": "taea",
				"icon": "https://img.esa.io/uploads/production/users/2/icon/thumb_m_2690997f07b7de3014a36d90827603d6.jpg"
			},
			"stargazers_count": 0,
			"star": false
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.GetComment(
				context.Background(),
				13,
			)
		},
	})
}

func TestClientImpl_CreateComment(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/posts/2/comments",
		body: `
		{"comment":{"body_md":"LGTM!"}}
		`,
		res: `
		{
			"id": 22767,
			"body_md": "LGTM!",
			"body_html": "<p>LGTM!</p>\n",
			"created_at": "2015-06-21T19:36:20+09:00",
			"updated_at": "2015-06-21T19:36:20+09:00",
			"url": "https://docs.esa.io/posts/2#comment-22767",
			"created_by": {
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
			},
			"stargazers_count": 0,
			"star": false
		}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			wrap := map[string]CreateCommentParam{}
			if err := json.Unmarshal([]byte(body), &wrap); err != nil {
				return nil, err
			}
			p := wrap["comment"]
			return client.CreateComment(
				context.Background(),
				2,
				p,
			)
		},
	})
}

func TestClientImpl_UpdateComment(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPatch,
		spath:  "/v1/teams/docs/comments/22767",
		body: `
		{"comment":{"body_md":"LGTM!!!"}}
		`,
		res: `
		{
			"id": 22767,
			"body_md": "LGTM! :sushi:",
			"body_html": "<p>LGTM!!!</p>\n",
			"created_at": "2015-06-21T19:36:20+09:00",
			"updated_at": "2015-06-21T19:40:33+09:00",
			"url": "https://docs.esa.io/posts/2#comment-22767",
			"created_by": {
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
			},
			"stargazers_count": 0,
			"star": false
		}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			wrap := map[string]UpdateCommentParam{}
			if err := json.Unmarshal([]byte(body), &wrap); err != nil {
				return nil, err
			}
			p := wrap["comment"]
			return client.UpdateComment(
				context.Background(),
				22767,
				p,
			)
		},
	})
}

func TestClientImpl_DeleteComment(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodDelete,
		spath:  "/v1/teams/docs/comments/22767",
		do: func(client Client) (interface{}, error) {
			return nil, client.DeleteComment(
				context.Background(),
				22767,
			)
		},
	})
}
