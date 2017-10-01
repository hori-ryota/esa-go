package esa

import (
	"context"
	"path"
	"time"
)

// CommentsResp is resp for comments
type CommentsResp struct {
	PageResp
	Comments []Comment `json:"comments"`
}

// Comment is struct for comment
type Comment struct {
	ID              uint      `json:"id"`
	BodyMd          string    `json:"body_md"`
	BodyHTML        string    `json:"body_html"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	URL             string    `json:"url"`
	CreatedBy       Member    `json:"created_by"`
	StargazersCount uint      `json:"stargazers_count"`
	Star            bool      `json:"star"`
}

// CreateCommentParam is param for create comment
type CreateCommentParam struct {
	BodyMD string  `json:"body_md"`
	User   *string `json:"user"`
}

// UpdateCommentParam is param for update comment
type UpdateCommentParam struct {
	BodyMD *string `json:"body_md"`
	User   *string `json:"user"`
}

// ListComments list comments
func (c ClientImpl) ListComments(ctx context.Context, postNumber uint, page uint, parPage uint) (*CommentsResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(postNumber), "comments")
	query := c.pagerQuery(page, parPage)
	res := CommentsResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetComment get comment
func (c ClientImpl) GetComment(ctx context.Context, id uint) (*Comment, error) {
	spath := path.Join("/v1/teams", c.teamName, "comments", uintToStr(id))
	res := Comment{}
	if err := c.httpGet(ctx, spath, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateComment create comment
func (c ClientImpl) CreateComment(ctx context.Context, postNumber uint, param CreateCommentParam) (*Comment, error) {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(postNumber), "comments")
	res := Comment{}
	if err := c.httpPost(ctx, spath, param, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateComment update comment
func (c ClientImpl) UpdateComment(ctx context.Context, id uint, param UpdateCommentParam) (*Comment, error) {
	spath := path.Join("/v1/teams", c.teamName, "comments", uintToStr(id))
	res := Comment{}
	if err := c.httpPatch(ctx, spath, param, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteComment delete comment
func (c ClientImpl) DeleteComment(ctx context.Context, id uint) error {
	spath := path.Join("/v1/teams", c.teamName, "comments", uintToStr(id))
	return c.httpDelete(ctx, spath)
}
