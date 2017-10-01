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

// InvitationByEmailParam is param for invitation by email
type InvitationByEmailParam struct {
	Member struct {
		Emails []string `json:"emails"`
	} `json:"member"`
}

// InvitationByEmailResp is resp for invitation by email
type InvitationByEmailResp struct {
	Invitations []Invitation `json:"invitations"`
}

// GetInvitationURL get invitationURL
func (c ClientImpl) GetInvitationURL(ctx context.Context) (*URLResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "invitation")
	res := URLResp{}
	if err := c.httpGet(ctx, spath, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RegenerateInvitationURL regenerate invitationURL
func (c ClientImpl) RegenerateInvitationURL(ctx context.Context) (*URLResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "invitation_regenerator")
	res := URLResp{}
	if err := c.httpPost(ctx, spath, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// InviteByEmail invite byEmail
func (c ClientImpl) InviteByEmail(ctx context.Context, emails ...string) (*InvitationByEmailResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "invitations")
	param := InvitationByEmailParam{}
	param.Member.Emails = emails
	res := InvitationByEmailResp{}
	if err := c.httpPost(ctx, spath, param, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ListInvitations list invitations
func (c ClientImpl) ListInvitations(ctx context.Context, page uint, perPage uint) (*InvitationsResp, error) {
	spath := path.Join("/v1/teams", c.teamName, "invitations")
	query := c.pagerQuery(page, perPage)
	res := InvitationsResp{}
	if err := c.httpGet(ctx, spath, query, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteInvitation delete invitation
func (c ClientImpl) DeleteInvitation(ctx context.Context, code string) error {
	spath := path.Join("/v1/teams", c.teamName, "invitations", code)
	return c.httpDelete(ctx, spath)
}
