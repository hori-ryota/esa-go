package esa

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/pkg/errors"
)

func (c ClientImpl) newRequest(ctx context.Context, method string, u url.URL, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to http.NewRequest")
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.userAgent())
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	return req, nil
}

func mergeQuery(values ...url.Values) url.Values {
	merged := url.Values{}
	for _, vs := range values {
		for k, v := range vs {
			merged[k] = append(merged[k], v...)
		}
	}
	return merged
}

func (c ClientImpl) pagerQuery(page uint, parPage uint) url.Values {
	v := url.Values{}
	v.Set("page", strconv.FormatUint(uint64(page), 10))
	v.Set("par_page", strconv.FormatUint(uint64(parPage), 10))
	return v
}

func (c ClientImpl) createURL(spath string, query url.Values) url.URL {
	u := c.baseURL
	u.Path = path.Join(u.Path, spath)
	values := mergeQuery(u.Query(), query)
	u.RawQuery = values.Encode()
	return u
}

func (c ClientImpl) decodeBody(resp *http.Response, res interface{}) error {
	defer resp.Body.Close()
	if res == nil {
		return nil
	}
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(res)
}

func (c ClientImpl) decodeError(resp *http.Response) (*Error, error) {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	e := Error{}
	if err := decoder.Decode(&e); err != nil {
		return nil, err
	}
	e.StatusCode = resp.StatusCode
	return &e, nil
}

func (c ClientImpl) parseResp(resp *http.Response, res interface{}) error {
	if resp.StatusCode >= 400 {
		e, err := c.decodeError(resp)
		if err != nil {
			return errors.Wrap(err, "failed to decode error")
		}
		return e
	}
	if err := c.decodeBody(resp, res); err != nil {
		return errors.Wrap(err, "failed to decode resp")
	}
	return nil
}

func (c ClientImpl) executeRequest(req *http.Request, res interface{}) error {
	resp, err := c.httpClient().Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to http request")
	}
	if err := c.parseResp(resp, res); err != nil {
		return errors.Wrap(err, "failed to parse resp")
	}
	return nil
}

func (c ClientImpl) httpWithQuery(ctx context.Context, method string, spath string, query url.Values, res interface{}) error {
	u := c.createURL(spath, query)
	req, err := c.newRequest(ctx, method, u, nil)
	if err != nil {
		return errors.Wrap(err, "failed to NewRequest")
	}
	if err := c.executeRequest(req, res); err != nil {
		return errors.Wrap(err, "failed to execute request")
	}
	return nil
}

func (c ClientImpl) httpGet(ctx context.Context, spath string, query url.Values, res interface{}) error {
	return c.httpWithQuery(ctx, http.MethodGet, spath, query, res)
}

func (c ClientImpl) httpWithBody(ctx context.Context, method string, spath string, body interface{}, res interface{}) error {
	if body == nil {
		return c.httpWithoutParam(ctx, method, spath, res)
	}
	u := c.createURL(spath, nil)
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(body); err != nil {
		return errors.Wrap(err, "failed to encode json")
	}
	req, err := c.newRequest(ctx, method, u, b)
	if err != nil {
		return errors.Wrap(err, "failed to NewRequest")
	}
	if err := c.executeRequest(req, res); err != nil {
		return errors.Wrap(err, "failed to execute request")
	}
	return nil
}

func (c ClientImpl) httpPost(ctx context.Context, spath string, body interface{}, res interface{}) error {
	return c.httpWithBody(ctx, http.MethodPost, spath, body, res)
}

func (c ClientImpl) httpPatch(ctx context.Context, spath string, body interface{}, res interface{}) error {
	return c.httpWithBody(ctx, http.MethodPatch, spath, body, res)
}

func (c ClientImpl) httpWithoutParam(ctx context.Context, method string, spath string, res interface{}) error {
	u := c.createURL(spath, nil)
	req, err := c.newRequest(ctx, method, u, nil)
	if err != nil {
		return errors.Wrap(err, "failed to NewRequest")
	}
	if err := c.executeRequest(req, res); err != nil {
		return errors.Wrap(err, "failed to execute request")
	}
	return nil
}

func (c ClientImpl) httpDelete(ctx context.Context, spath string) error {
	return c.httpWithoutParam(ctx, http.MethodPatch, spath, nil)
}
