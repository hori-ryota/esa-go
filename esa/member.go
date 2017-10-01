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
	Name       string `json:"name,omitempty"`
	ScreenName string `json:"screen_name,omitempty"`
	Icon       string `json:"icon,omitempty"`
	Email      string `json:"email,omitempty"`
	PostsCount int64  `json:"posts_count,omitempty"`
}

// ListMembers list members
func (c ClientImpl) ListMembers(ctx context.Context, page uint, perPage uint) (*MembersResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "members")
	query := c.pagerQuery(page, perPage)
	res := MembersResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
