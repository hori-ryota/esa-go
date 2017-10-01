package esa

import (
	"context"
	"path"
)

// BatchMoveParam is param for batch move
type BatchMoveParam struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// BatchMoveResult is result for batch move
type BatchMoveResult struct {
	Count uint   `json:"count"`
	From  string `json:"from"`
	To    string `json:"to"`
}

// ChangeCategory change category
func (c ClientImpl) ChangeCategory(ctx context.Context, from string, to string) (*BatchMoveResult, error) {
	param := BatchMoveParam{
		From: from,
		To:   to,
	}
	spath := path.Join("/v1/teams", c.teamName, "categories", "batch_move")
	res := BatchMoveResult{}
	if err := c.httpPost(ctx, spath, param, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
