package esa

import (
	"context"
	"path"
	"time"
)

// StargazersResp is resp for stargazers
type StargazersResp struct {
	PageResp
	Stargazers []Stargazer `json:"stargazers"`
}

// Stargazer is struct for stargazer
type Stargazer struct {
	CreatedAt time.Time `json:"created_at"`
	Body      *string   `json:"body"`
	User      Member    `json:"user"`
}

// StarParam is param for create star
type StarParam struct {
	Body *string `json:"body"`
}

// ListPostStargazers list postStargazers
func (c ClientImpl) ListPostStargazers(ctx context.Context, postNumber uint, page uint, perPage uint) (*StargazersResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(postNumber), "stargazers")
	query := c.pagerQuery(page, perPage)
	res := StargazersResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// StarPost star post
func (c ClientImpl) StarPost(ctx context.Context, postNumber uint, param StarParam) error {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(postNumber), "star")
	return c.httpPost(ctx, spath, param, nil)
}

// UnstarPost unstar post
func (c ClientImpl) UnstarPost(ctx context.Context, postNumber uint) error {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(postNumber), "star")
	return c.httpDelete(ctx, spath)
}

// ListCommentStargazers list commentStargazers
func (c ClientImpl) ListCommentStargazers(ctx context.Context, commentID uint, page uint, perPage uint) (*StargazersResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "comments", uintToStr(commentID), "stargazers")
	query := c.pagerQuery(page, perPage)
	res := StargazersResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// StarComment star comment
func (c ClientImpl) StarComment(ctx context.Context, commentID uint, param StarParam) error {
	spath := path.Join("/v1/teams", c.teamName, "comments", uintToStr(commentID), "star")
	return c.httpPost(ctx, spath, param, nil)
}

// UnstarComment unstar comment
func (c ClientImpl) UnstarComment(ctx context.Context, commentID uint) error {
	spath := path.Join("/v1/teams", c.teamName, "comments", uintToStr(commentID), "star")
	return c.httpDelete(ctx, spath)
}
