package esa

import (
	"context"
	"path"
)

// MembersResp is resp for members
type MembersResp struct {
	PageResp
	Members []Member `json:"members"`
}

// Member is struct for member
type Member struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
	Email      string `json:"email"`
	PostsCount int64  `json:"posts_count"`
}

// ListMembers list members
func (c ClientImpl) ListMembers(ctx context.Context, page uint, parPage uint) (*MembersResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "members")
	query := c.pagerQuery(page, parPage)
	res := MembersResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
