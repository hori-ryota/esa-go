package esa

import (
	"context"
	"net/http"
	"testing"
)

func TestClientImpl_ListMembers(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/members",
		query:  "page=2&per_page=30",
		res: `
		{
			"members": [
			{
				"name": "Atsuo Fukaya",
				"screen_name": "fukayatsu",
				"icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png",
				"email": "fukayatsu@esa.io",
				"posts_count": 222
			},
			{
				"name": "TAEKO AKATSUKA",
				"screen_name": "taea",
				"icon": "https://img.esa.io/uploads/production/users/2/icon/thumb_m_2690997f07b7de3014a36d90827603d6.jpg",
				"email": "taea@esa.io",
				"posts_count": 111
			}
			],
			"prev_page": null,
			"next_page": null,
			"total_count": 2,
			"page": 1,
			"per_page": 20,
			"max_per_page": 100
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.ListMembers(
				context.Background(),
				2,
				30,
			)
		},
	})
}
