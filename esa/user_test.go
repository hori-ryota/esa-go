package esa

import (
	"context"
	"net/http"
	"testing"
)

func TestClientImpl_GetUser(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/user",
		query:  "include=teams",
		res: `
		{
			"id": 1,
			"name": "Atsuo Fukaya",
			"screen_name": "fukayatsu",
			"created_at": "2014-05-10T11:50:07+09:00",
			"updated_at": "2016-04-17T12:35:16+09:00",
			"icon": "https://img.esa.io/uploads/production/users/1/icon/thumb_m_402685a258cf2a33c1d6c13a89adec92.png",
			"email": "fukayatsu@esa.io"
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.GetUser(
				context.Background(),
				GetUserParam{
					Include: []GetUserParamInclude{
						GetUserParamIncludeTeams,
					},
				},
			)
		},
	})
}
