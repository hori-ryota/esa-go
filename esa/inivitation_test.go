package esa

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestClientImpl_GetInvitationURL(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/invitation",
		res: `
		{
			"url": "https://docs.esa.io/team/invitations/member-c05d112fa34870998ab4da1e98846ae3"
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.GetInvitationURL(
				context.Background(),
			)
		},
	})
}

func TestClientImpl_RegenerateInvitationURL(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/invitation_regenerator",
		res: `
		{
			"url": "https://docs.esa.io/team/invitations/member-58891f72edcbb8ac22f1e5548b0128d9"
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.RegenerateInvitationURL(
				context.Background(),
			)
		},
	})
}

func TestClientImpl_InviteByEmail(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodPost,
		spath:  "/v1/teams/docs/invitations",
		body: `
		{
			"member": {
				"emails": ["foo@example.com", "bar@example.com"]
			}
		}
		`,
		res: `
		{
			"invitations": [
			{
				"email": "foo@example.com",
				"code": "mee93383edf699b525e01842d34078e28",
				"expires_at": "2017-08-17T12:00:41+09:00",
				"url": "https://docs.esa.io/team/invitations/mee93383edf699b525e01842d34078e28/join"
			},
			{
				"email": "bar@example.com",
				"code": "m934f1f60732f49d50ee5b3f96841ff13",
				"expires_at": "2017-08-17T12:00:41+09:00",
				"url": "https://docs.esa.io/team/invitations/m934f1f60732f49d50ee5b3f96841ff13/join"
			}
			]
		}
		`,
		doWithBody: func(client Client, body string) (interface{}, error) {
			p := InvitationByEmailParam{}
			if err := json.Unmarshal([]byte(body), &p); err != nil {
				return nil, err
			}
			return client.InviteByEmail(
				context.Background(),
				p.Member.Emails...,
			)
		},
	})
}

func TestClientImpl_ListInvitations(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodGet,
		spath:  "/v1/teams/docs/invitations",
		query:  "page=2&per_page=30",
		res: `
		{
			"invitations": [
			{
				"email": "foo@example.com",
				"code": "mee93383edf699b525e01842d34078e28",
				"expires_at": "2017-08-17T12:00:41+09:00",
				"url": "https://docs.esa.io/team/invitations/mee93383edf699b525e01842d34078e28/join"
			},
			{
				"email": "bar@example.com",
				"code": "mc542eed211a8e4f1db6ccccb14fcda9d",
				"expires_at": "2017-08-17T12:00:44+09:00",
				"url": "https://docs.esa.io/team/invitations/mc542eed211a8e4f1db6ccccb14fcda9d/join"
			}
			],
			"prev_page": null,
			"next_page": null,
			"total_count": 2,
			"page": 1,
			"per_page": 20,
			"max_per_page": 100
		}
		`,
		do: func(client Client) (interface{}, error) {
			return client.ListInvitations(
				context.Background(),
				2,
				30,
			)
		},
	})
}

func TestClientImpl_DeleteInvitation(t *testing.T) {
	httpTest(t, httpTestCase{
		method: http.MethodDelete,
		spath:  "/v1/teams/docs/invitations/mee93383edf699b525e01842d34078e28",
		do: func(client Client) (interface{}, error) {
			return nil, client.DeleteInvitation(
				context.Background(),
				"mee93383edf699b525e01842d34078e28",
			)
		},
	})
}
