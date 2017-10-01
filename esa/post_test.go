package esa

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestClientImpl_ListPosts(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/posts",
		query:  "page=2&per_page=30&q=query&include=stargazers,comments&order=asc",
		res: `
		{
			"posts": [
			{
				"number": 1,
				"name": "hi!",
				"full_name": "日報/2015/05/09/hi! #api #dev",
				"wip": true,
				"body_md": "# Getting Started",
				"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
				"created_at": "2015-05-09T11:54:50+09:00",
				"message": "Add Getting Started section",
				"url": "https://docs.esa.io/posts/1",
				"updated_at": "2015-05-09T11:54:51+09:00",
				"tags": [
				"api",
				"dev"
				],
				"category": "日報/2015/05/09",
				"revision_number": 1,
				"created_by": {
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
				},
				"updated_by": {
					"name": "Atsuo Fukaya",
					"screen_name": "fukayatsu",
					"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
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
			return client.ListPosts(
				context.Background(),
				ListPostsParam{
					Q: "query",
					Include: []ListPostsParamInclude{
						ListPostsParamIncludeStargazers,
						ListPostsParamIncludeComments,
					},
					Order: ASC,
				},
				2,
				30,
			)
		},
	})
}

func TestClientImpl_GetPost(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/posts/1",
		res: `
		{
			"number": 1,
			"name": "hi!",
			"full_name": "日報/2015/05/09/hi! #api #dev",
			"wip": true,
			"body_md": "# Getting Started",
			"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
			"created_at": "2015-05-09T11:54:50+09:00",
			"message": "Add Getting Started section",
			"url": "https://docs.esa.io/posts/1",
			"updated_at": "2015-05-09T11:54:51+09:00",
			"tags": [
			"api",
			"dev"
			],
			"category": "日報/2015/05/09",
			"revision_number": 1,
			"created_by": {
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
			},
			"updated_by": {
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
			},
			"kind": "flow",
			"comments_count": 1,
			"tasks_count": 1,
			"done_tasks_count": 1,
			"stargazers_count": 1,
			"watchers_count": 1,
			"star": true,
			"watch": true
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.GetPost(
				context.Background(),
				1,
			)
		},
	})
}

func TestClientImpl_CreatePost(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/posts",
		body: `
		{
			"post":{
				"name":"hi!",
				"body_md":"# Getting Started\n",
				"tags":[
				"api",
				"dev"
				],
				"category":"dev/2015/05/10",
				"wip":false,
				"message":"Add Getting Started section"
			}
		}
		`,
		res: `
		{
			"number": 5,
			"name": "hi!",
			"full_name": "dev/2015/05/10/hi! #api #dev",
			"wip": false,
			"body_md": "# Getting Started\n",
			"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
			"created_at": "2015-05-09T12:12:37+09:00",
			"message": "Add Getting Started section",
			"url": "https://docs.esa.io/posts/5",
			"updated_at": "2015-05-09T12:12:37+09:00",
			"tags": [
			"api",
			"dev"
			],
			"category": "dev/2015/05/10",
			"revision_number": 1,
			"created_by": {
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
			},
			"updated_by": {
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
			},
			"kind": "flow",
			"comments_count": 0,
			"tasks_count": 0,
			"done_tasks_count": 0,
			"stargazers_count": 0,
			"watchers_count": 1,
			"star": false,
			"watch": false
		}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			wrap := map[string]CreatePostParam{}
			if err := json.Unmarshal([]byte(body), &wrap); err != nil {
				return nil, err
			}
			p := wrap["post"]
			return client.CreatePost(
				context.Background(),
				p,
			)
		},
	})
}

func TestClientImpl_UpdatePost(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPatch,
		spath:  "/v1/teams/docs/posts/5",
		body: `
		{
			"post":{
				"name":"hi!",
				"body_md":"# Getting Started\n",
				"tags":[
				"api",
				"dev"
				],
				"category":"dev/2015/05/10",
				"wip":false,
				"message":"Add Getting Started section",
				"original_revision": {
					"body_md": "# Getting ...",
					"number":1,
					"user": "fukayatsu"
				}
			}
		}
		`,
		res: `
		{
			"number": 5,
			"name": "hi!",
			"full_name": "日報/2015/05/10/hi! #api #dev",
			"wip": false,
			"body_md": "# Getting Started\n",
			"body_html": "<h1 id=\"1-0-0\" name=\"1-0-0\">\n<a class=\"anchor\" href=\"#1-0-0\"><i class=\"fa fa-link\"></i><span class=\"hidden\" data-text=\"Getting Started\"> &gt; Getting Started</span></a>Getting Started</h1>\n",
			"created_at": "2015-05-09T12:12:37+09:00",
			"message": "Add Getting Started section",
			"url": "https://docs.esa.io/posts/5",
			"updated_at": "2015-05-09T12:19:48+09:00",
			"tags": [
			"api",
			"dev"
			],
			"category": "日報/2015/05/10",
			"revision_number": 2,
			"created_by": {
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
			},
			"updated_by": {
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "http://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png"
			},
			"overlapped": false,
			"kind": "flow",
			"comments_count": 0,
			"tasks_count": 0,
			"done_tasks_count": 0,
			"stargazers_count": 0,
			"watchers_count": 1,
			"star": false,
			"watch": false
		}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			wrap := map[string]UpdatePostParam{}
			if err := json.Unmarshal([]byte(body), &wrap); err != nil {
				return nil, err
			}
			p := wrap["post"]
			return client.UpdatePost(
				context.Background(),
				5,
				p,
			)
		},
	})
}

func TestClientImpl_DeletePost(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodDelete,
		spath:  "/v1/teams/docs/posts/5",
		do: func(client Client) (interface{}, error) {
			return nil, client.DeletePost(
				context.Background(),
				5,
			)
		},
	})
}
