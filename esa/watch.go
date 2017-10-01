package esa

import (
	"context"
	"path"
	"time"
)

// WatchersResp is resp for watchers
type WatchersResp struct {
	PageResp
	Watchers []Watcher `json:"watchers"`
}

// Watcher is struct for watcher
type Watcher struct {
	CreatedAt time.Time `json:"created_at"`
	User      Member    `json:"user"`
}

// ListPostWatchers list postWatchers
func (c ClientImpl) ListPostWatchers(ctx context.Context, postNumber uint, page uint, parPage uint) (*WatchersResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(postNumber), "watchers")
	query := c.pagerQuery(page, parPage)
	res := WatchersResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// WatchPost watch post
func (c ClientImpl) WatchPost(ctx context.Context, postNumber uint) error {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(postNumber), "watch")
	return c.httpPost(ctx, spath, nil, nil)
}

// UnwatchPost unwatch post
func (c ClientImpl) UnwatchPost(ctx context.Context, postNumber uint) error {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(postNumber), "watch")
	return c.httpDelete(ctx, spath)
}
