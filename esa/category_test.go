package esa

import (
	"context"
	"net/http"
	"testing"
)

func TestClientImpl_ChangeCategory(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/categories/batch_move",
		body: `
		{
			"from": "/foo/bar/",
			"to": "/baz/"
		}
		`,
		res: `
		{
			"count": 3,
			"from": "/foo/bar/",
			"to": "/baz/"
		}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			return client.ChangeCategory(
				context.Background(),
				"/foo/bar/",
				"/baz/",
			)
		},
	})
}
