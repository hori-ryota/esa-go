package esa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

type httpTestCase struct {
	do         func(client Client) (res interface{}, err error)
	doWithBody func(client Client, body string) (res interface{}, err error)

	method string
	spath  string
	// optional
	query string
	// optional
	body string
	// optional
	res string
	// optional
	statusCode int
	// optional. default:doc
	teamName string
	// optional
	wantError bool
}

// httpTest is test util for common request
func httpTest(
	t *testing.T,
	tt httpTestCase,
) {
	assert := require.New(t)

	teamName := "docs"
	if tt.teamName != "" {
		teamName = tt.teamName
	}

	accessToken := "accessToken"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(tt.method, r.Method)

		assert.Equal(tt.spath, r.URL.Path)

		ttQuery, err := url.ParseQuery(tt.query)
		assert.NoError(err)
		assert.Equal(
			ttQuery,
			r.URL.Query(),
		)

		b, err := ioutil.ReadAll(r.Body)
		assert.NoError(err)
		if tt.body != "" {
			assert.JSONEq(tt.body, string(b))
		} else {
			assert.Empty(b)
		}

		_, err = w.Write([]byte(tt.res))
		assert.NoError(err)
		if tt.statusCode >= 400 {
			http.Error(w, tt.res, tt.statusCode)
			return
		}
		if tt.statusCode != 0 {
			w.WriteHeader(tt.statusCode)
		}
	}))
	defer ts.Close()

	client := NewClient(accessToken, teamName)
	u, err := url.Parse(ts.URL)
	assert.NoError(err)
	client.OverwriteBaseURL(*u)

	var res interface{}
	if tt.do != nil {
		res, err = tt.do(client)
	} else {
		res, err = tt.doWithBody(client, tt.body)
	}
	if tt.wantError {
		assert.Error(err)
		return
	}
	assert.NoError(err)
	if res == nil {
		return
	}
	j, err := json.Marshal(res)
	assert.NoError(err)
	assert.JSONEq(tt.res, string(j))
}
