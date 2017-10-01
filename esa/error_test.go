package esa

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestClientImpl_Error(t *testing.T) {
	assert := require.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(
			w,
			`
			{
				"error": "not_found",
				"message": "Not found"
			}
			`,
			http.StatusBadRequest,
		)
	}))

	client := NewClient("", "")
	u, err := url.Parse(ts.URL)
	assert.NoError(err)
	client.OverwriteBaseURL(*u)

	_, err = client.GetTeam(context.Background())
	assert.EqualError(
		errors.Cause(err),
		Error{
			Err:        "not_found",
			Message:    "Not found",
			StatusCode: http.StatusBadRequest,
		}.Error(),
	)
}
