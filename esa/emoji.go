package esa

import (
	"context"
	"path"

	queryPkg "github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// EmojisResp is resp for emojis
type EmojisResp struct {
	PageResp
	Emojis []Emoji `json:"emojis"`
}

// Emoji is struct for emoji
type Emoji struct {
	Code    string   `json:"code"`
	Aliases []string `json:"aliases"`
	URL     string   `json:"url"`
}

// ListEmojisParamInclude is enum for post param "include"
type ListEmojisParamInclude string

const (
	// ListEmojisParamIncludeAll is enum for post param "include"
	ListEmojisParamIncludeAll ListEmojisParamInclude = "all"
)

// ListEmojisParam is param for fetch user
type ListEmojisParam struct {
	Include []ListEmojisParamInclude `url:"include,comma"`
}

// CreateEmojiParam is param for create emoji
type CreateEmojiParam struct {
	Code string `json:"code"`
	// For alias
	OriginCode *string `json:"origin_code"`
	// BASE64 String
	Image *string `json:"image"`
}

// ListEmojis list emojis
func (c ClientImpl) ListEmojis(ctx context.Context, param ListEmojisParam, page uint, parPage uint) (*EmojisResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "emojis")
	pagerQuery := c.pagerQuery(page, parPage)
	query, err := queryPkg.Values(param)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create query from param")
	}
	query = mergeQuery(query, pagerQuery)
	res := EmojisResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateEmoji create emoji
func (c ClientImpl) CreateEmoji(ctx context.Context, param CreateEmojiParam) (*Emoji, error) {
	spath := path.Join("/v1/teams", c.teamName, "emojis")
	res := Emoji{}
	wrap := wrapRes("emoji", &res)
	if err := c.httpPost(ctx, spath, param, wrap); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteEmoji delete emoji
func (c ClientImpl) DeleteEmoji(ctx context.Context, code string) error {
	spath := path.Join("/v1/teams", c.teamName, "emojis", code)
	return c.httpDelete(ctx, spath)
}
