package esa

import (
	"context"
	"path"
	"time"
)

// InvitationsResp is resp for invitations
type InvitationsResp struct {
	PageResp
	Invitations []Invitation `json:"invitations"`
}

// Invitation is struct for inivitaion
type Invitation struct {
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
	URL       string    `json:"url"`
}

// InvitationByEmailParam is param for invitaion by email
type InvitationByEmailParam struct {
	Member struct {
		Emails []string `json:"emails"`
	} `json:"member"`
}

// GetInvitationURL get invitationURL
func (c ClientImpl) GetInvitationURL(ctx context.Context) (*URLResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "invitaion")
	res := URLResp{}
	if err := c.httpGet(ctx, spath, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RegenerateInvitationURL regenerate invitationURL
func (c ClientImpl) RegenerateInvitationURL(ctx context.Context) (*URLResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "invitaion_regenerator")
	res := URLResp{}
	if err := c.httpPost(ctx, spath, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// InviteByEmail invite byEmail
func (c ClientImpl) InviteByEmail(ctx context.Context, emails ...string) (*InvitationsResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "invitaions")
	param := InvitationByEmailParam{}
	param.Member.Emails = emails
	res := InvitationsResp{}
	if err := c.httpPost(ctx, spath, param, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListInvitations list invitations
func (c ClientImpl) ListInvitations(ctx context.Context, page uint, parPage uint) (*InvitationsResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "invitaions")
	query := c.pagerQuery(page, parPage)
	res := InvitationsResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteInvitation delete invitation
func (c ClientImpl) DeleteInvitation(ctx context.Context, code string) error {
	spath := path.Join("/v1/teams", c.teamName, "invitaions", code)
	return c.httpDelete(ctx, spath)
}
