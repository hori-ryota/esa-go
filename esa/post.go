package esa

import (
	"context"
	"path"
	"time"

	queryPkg "github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// PostsResp is resp for posts
type PostsResp struct {
	PageResp
	Posts []Post `json:"posts"`
}

// PostKind is enum for post kind
type PostKind string

const (
	// PostKindStock is enum for post kind
	PostKindStock PostKind = "stock"
	// PostKindFlow is enum for post kind
	PostKindFlow PostKind = "flow"
)

// Post is struct for post
type Post struct {
	Name            string    `json:"name"`
	Number          uint      `json:"number"`
	Tags            []string  `json:"tags"`
	Category        string    `json:"category"`
	FullName        string    `json:"full_name"`
	WIP             bool      `json:"wip"`
	BodyMD          string    `json:"body_md"`
	BodyHTML        string    `json:"body_html"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Message         string    `json:"message"`
	RevisionNumber  int64     `json:"revision_number"`
	CreatedBy       Member    `json:"created_by"`
	UpdatedBy       Member    `json:"updated_by"`
	Kind            PostKind  `json:"kind"`
	URL             string    `json:"url"`
	CommentsCount   uint      `json:"comments_count"`
	TasksCount      uint      `json:"tasks_count"`
	DoneTasksCount  uint      `json:"done_tasks_count"`
	StargazersCount uint      `json:"stargazers_count"`
	WatchersCount   uint      `json:"watchers_count"`
	Star            bool      `json:"star"`
	Watch           bool      `json:"watch"`
}

// ListPostsParamInclude is enum for post param "include"
type ListPostsParamInclude string

const (
	// ListPostsParamIncludeComments is enum for post param "include"
	ListPostsParamIncludeComments ListPostsParamInclude = "comments"
	// ListPostsParamIncludeCommentsStargazers is enum for post param "include"
	ListPostsParamIncludeCommentsStargazers ListPostsParamInclude = "comments.stargazers"
	// ListPostsParamIncludeStargazers is enum for post param "include"
	ListPostsParamIncludeStargazers ListPostsParamInclude = "stargazers"
)

// ListPostsParamSort is enum for post param "sort"
type ListPostsParamSort string

const (
	// ListPostsParamSortUpdated is enum for post param "sort"
	ListPostsParamSortUpdated ListPostsParamSort = "updated"
	// ListPostsParamSortCreated is enum for post param "sort"
	ListPostsParamSortCreated ListPostsParamSort = "created"
	// ListPostsParamSortStars is enum for post param "sort"
	ListPostsParamSortStars ListPostsParamSort = "stars"
	// ListPostsParamSortWatches is enum for post param "sort"
	ListPostsParamSortWatches ListPostsParamSort = "watches"
	// ListPostsParamSortComments is enum for post param "sort"
	ListPostsParamSortComments ListPostsParamSort = "comments"
	// ListPostsParamSortBestMatch is enum for post param "sort"
	ListPostsParamSortBestMatch ListPostsParamSort = "best_match"
)

// ListPostsParam is param for list posts
type ListPostsParam struct {
	Q       string                  `url:"q"`
	Include []ListPostsParamInclude `url:"include,comma"`
	Sort    ListPostsParamSort      `url:"sort"`
	Order   Order                   `url:"order"`
}

// CreatePostParam is param for create post
type CreatePostParam struct {
	Name     string   `json:"name"`
	BodyMD   string   `json:"body_md"`
	Tags     []string `json:"tags"`
	Category string   `json:"category"`
	WIP      bool     `json:"wip"`
	Message  string   `json:"message"`
	User     *string  `json:"user"`
}

// UpdatePostParam is param for update post
type UpdatePostParam struct {
	Name             *string   `json:"name"`
	BodyMD           *string   `json:"body_md"`
	Tags             *[]string `json:"tags"`
	Category         *string   `json:"category"`
	WIP              *bool     `json:"wip"`
	Message          *string   `json:"message"`
	CreatedBy        *string   `json:"created_by"`
	UpdatedBy        *string   `json:"updated_by"`
	OriginalRevision struct {
		BodyMD string `json:"body_md"`
		Number uint   `json:"number"`
		User   string `json:"user"`
	} `json:"original_revision"`
}

// ListPosts list posts
func (c ClientImpl) ListPosts(ctx context.Context, param ListPostsParam, page uint, parPage uint) (*PostsResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "posts")
	pagerQuery := c.pagerQuery(page, parPage)
	query, err := queryPkg.Values(param)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create query from param")
	}
	query = mergeQuery(query, pagerQuery)
	res := PostsResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetPost get post
func (c ClientImpl) GetPost(ctx context.Context, number uint) (*Post, error) {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(number))
	res := Post{}
	if err := c.httpGet(ctx, spath, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreatePost create post
func (c ClientImpl) CreatePost(ctx context.Context, param CreatePostParam) (*Post, error) {
	spath := path.Join("/v1/teams", c.teamName, "posts")
	res := Post{}
	wrap := wrapRes("post", &res)
	if err := c.httpPost(ctx, spath, param, wrap); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdatePost update post
func (c ClientImpl) UpdatePost(ctx context.Context, number uint, param UpdatePostParam) (*Post, error) {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(number))
	res := Post{}
	wrap := wrapRes("post", &res)
	if err := c.httpPatch(ctx, spath, param, wrap); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeletePost delete post
func (c ClientImpl) DeletePost(ctx context.Context, number uint) error {
	spath := path.Join("/v1/teams", c.teamName, "posts", uintToStr(number))
	return c.httpDelete(ctx, spath)
}
