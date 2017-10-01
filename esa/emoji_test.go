package esa

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestClientImpl_ListEmojis(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/emojis",
		query:  "page=2&per_page=30&include=all",
		res: `
		{
			"emojis": [
			{
				"code": "grinning",
				"aliases": [
				"grinning"
				],
				"url": "https://assets.esa.io/images/emoji/unicode/1f600.png"
			},
			{
				"code": "smiley",
				"aliases": [
				"smiley"
				],
				"url": "https://assets.esa.io/images/emoji/unicode/1f603.png"
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
			return client.ListEmojis(
				context.Background(),
				ListEmojisParam{
					Include: []ListEmojisParamInclude{
						ListEmojisParamIncludeAll,
					},
				},
				2,
				30,
			)
		},
	})
}

func TestClientImpl_CreateEmoji(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/emojis",
		body: `
		{
			"emoji": {
				"code": "team_emoji",
				"image": "dummy"
			}
		}
		`,
		res: `
		{
			"code": "team_emoji"
		}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			wrap := map[string]CreateEmojiParam{}
			if err := json.Unmarshal([]byte(body), &wrap); err != nil {
				return nil, err
			}
			p := wrap["emoji"]
			return client.CreateEmoji(
				context.Background(),
				p,
			)
		},
	})
}

func TestClientImpl_DeleteEmoji(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodDelete,
		spath:  "/v1/teams/docs/emojis/team_emoji",
		do: func(client Client) (interface{}, error) {
			return nil, client.DeleteEmoji(
				context.Background(),
				"team_emoji",
			)
		},
	})
}
