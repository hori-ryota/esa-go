package esa

import (
	"context"
	"path"
)

// TeamsResp is resp for teams
type TeamsResp struct {
	PageResp
	Teams []Team `json:"teams"`
}

// Privacy is enum for team privacy
type Privacy string

const (
	// PrivacyClosed is enum for team privacy
	PrivacyClosed Privacy = "closed"
	// PrivacyOpen is enum for team privacy
	PrivacyOpen Privacy = "open"
)

// Team is struct for team
type Team struct {
	Name        string  `json:"name"`
	Privacy     Privacy `json:"privacy"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	URL         string  `json:"url"`
}

// TeamStats is struct for team status
type TeamStats struct {
	Members            uint `json:"members"`
	Posts              uint `json:"posts"`
	PostsWIP           uint `json:"posts_wip"`
	PostsShipped       uint `json:"posts_shipped"`
	Comments           uint `json:"comments"`
	Stars              uint `json:"stars"`
	DailyActiveUsers   uint `json:"daily_active_users"`
	WeeklyActiveUsers  uint `json:"weekly_active_users"`
	MonthlyActiveUsers uint `json:"monthly_active_users"`
}

// ListTeams list teams
func (c ClientImpl) ListTeams(ctx context.Context, page uint, parPage uint) (*TeamsResp, error) {
	spath := "/v1/teams"
	query := c.pagerQuery(page, parPage)
	res := TeamsResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetTeam get team
func (c ClientImpl) GetTeam(ctx context.Context) (*Team, error) {
	spath := path.Join("/v1/teams", c.teamName)
	res := Team{}
	if err := c.httpGet(ctx, spath, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetTeamStats get teamStats
func (c ClientImpl) GetTeamStats(ctx context.Context) (*TeamStats, error) {
	spath := path.Join("/v1/teams", c.teamName, "stats")
	res := TeamStats{}
	if err := c.httpGet(ctx, spath, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
