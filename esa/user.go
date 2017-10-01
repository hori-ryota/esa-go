package esa

import (
	"context"
	"time"

	queryPkg "github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// User is struct for user
type User struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	ScreenName string    `json:"screen_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Icon       string    `json:"icon"`
	Email      string    `json:"email"`
}

// GetUserParamInclude is enum for post param "include"
type GetUserParamInclude string

const (
	// GetUserParamIncludeTeams is enum for post param "include"
	GetUserParamIncludeTeams GetUserParamInclude = "teams"
)

// GetUserParam is param for get user
type GetUserParam struct {
	Include []GetUserParamInclude `url:"include,comma"`
}

// GetUser get user
func (c ClientImpl) GetUser(ctx context.Context, param GetUserParam) (*User, error) {
	spath := "/v1/user"
	query, err := queryPkg.Values(param)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create query from param")
	}
	res := User{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
