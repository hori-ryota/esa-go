package esa

import (
	"context"
	"net/http"
	"testing"
)

func TestClientImpl_ListTeams(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams",
		query:  "page=2&per_page=30",
		res: `
		{
			"teams": [
			{
				"name": "docs",
				"privacy": "open",
				"description": "esa.io official documents",
				"icon": "https://img.esa.io/uploads/production/teams/105/icon/thumb_m_0537ab827c4b0c18b60af6cdd94f239c.png",
				"url": "https://docs.esa.io/"
			}
			],
			"prev_page": null,
			"next_page": null,
			"total_count": 1,
			"page": 1,
			"per_page": 20,
			"max_per_page": 100
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.ListTeams(
				context.Background(),
				2,
				30,
			)
		},
	})
}

func TestClientImpl_GetTeam(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs",
		res: `
		{
			"name": "docs",
			"privacy": "open",
			"description": "esa.io official documents",
			"icon": "https://img.esa.io/uploads/production/teams/105/icon/thumb_m_0537ab827c4b0c18b60af6cdd94f239c.png",
			"url": "https://docs.esa.io/"
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.GetTeam(
				context.Background(),
			)
		},
	})
}

func TestClientImpl_GetTeamStats(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/stats",
		res: `
		{
			"members": 20,
			"posts": 1959,
			"posts_wip": 59,
			"posts_shipped": 1900,
			"comments": 2695,
			"stars": 3115,
			"daily_active_users": 8,
			"weekly_active_users": 14,
			"monthly_active_users": 15
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.GetTeamStats(
				context.Background(),
			)
		},
	})
}
