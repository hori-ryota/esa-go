package esa

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
)

// Client is interface for esa
type Client interface {
	SetTeam(teamName string)
	SetAccessToken(accessToken string)
	AppendOption(options ...BaseOption)
	SetOption(options ...BaseOption)
	ClearOptions()
	// like url.URL{Scheme: "http", Host: "localhost"}
	OverwriteBaseURL(u url.URL)

	ListTeams(ctx context.Context, page uint, perPage uint) (*TeamsResp, error)
	GetTeam(ctx context.Context) (*Team, error)
	GetTeamStats(ctx context.Context) (*TeamStats, error)

	ListMembers(ctx context.Context, page uint, perPage uint) (*MembersResp, error)

	ListPosts(ctx context.Context, param ListPostsParam, page uint, perPage uint) (*PostsResp, error)
	GetPost(ctx context.Context, number uint) (*Post, error)
	CreatePost(ctx context.Context, param CreatePostParam) (*Post, error)
	UpdatePost(ctx context.Context, number uint, param UpdatePostParam) (*UpdatedPost, error)
	DeletePost(ctx context.Context, number uint) error

	ListComments(ctx context.Context, postNumber uint, page uint, perPage uint) (*CommentsResp, error)
	GetComment(ctx context.Context, id uint) (*Comment, error)
	CreateComment(ctx context.Context, postNumber uint, param CreateCommentParam) (*Comment, error)
	UpdateComment(ctx context.Context, id uint, param UpdateCommentParam) (*Comment, error)
	DeleteComment(ctx context.Context, id uint) error

	ListPostStargazers(ctx context.Context, postNumber uint, page uint, perPage uint) (*StargazersResp, error)
	StarPost(ctx context.Context, postNumber uint, param StarParam) error
	UnstarPost(ctx context.Context, postNumber uint) error

	ListCommentStargazers(ctx context.Context, commentID uint, page uint, perPage uint) (*StargazersResp, error)
	StarComment(ctx context.Context, commentID uint, param StarParam) error
	UnstarComment(ctx context.Context, commentID uint) error

	ListPostWatchers(ctx context.Context, postNumber uint, page uint, perPage uint) (*WatchersResp, error)
	WatchPost(ctx context.Context, postNumber uint) error
	UnwatchPost(ctx context.Context, postNumber uint) error

	ChangeCategory(ctx context.Context, from string, to string) (*BatchMoveResult, error)

	GetInvitationURL(ctx context.Context) (*URLResp, error)
	RegenerateInvitationURL(ctx context.Context) (*URLResp, error)
	InviteByEmail(ctx context.Context, emails ...string) (*InvitationByEmailResp, error)
	ListInvitations(ctx context.Context, page uint, perPage uint) (*InvitationsResp, error)
	DeleteInvitation(ctx context.Context, code string) error

	ListEmojis(ctx context.Context, param ListEmojisParam, page uint, perPage uint) (*EmojisResp, error)
	CreateEmoji(ctx context.Context, param CreateEmojiParam) (*Emoji, error)
	DeleteEmoji(ctx context.Context, code string) error

	GetUser(ctx context.Context, param GetUserParam) (*User, error)
}

type baseOptions struct {
	userAgent  string
	httpClient *http.Client
}

// BaseOption overwrite base requrest params
type BaseOption func(*baseOptions)

// WithUserAgent apply overwrite userAgent
func WithUserAgent(userAgent string) BaseOption {
	return func(ops *baseOptions) {
		ops.userAgent = userAgent
	}
}

// WithHTTPClient apply overwrite httpClient
func WithHTTPClient(httpClient *http.Client) BaseOption {
	return func(ops *baseOptions) {
		ops.httpClient = httpClient
	}
}

// ClientImpl is implementation of Client
type ClientImpl struct {
	accessToken string
	teamName    string
	options     baseOptions
	baseURL     url.URL
}

// NewClient initialize Client
func NewClient(accessToken string, teamName string, options ...BaseOption) Client {
	ops := baseOptions{}
	for _, o := range options {
		o(&ops)
	}
	baseURL := url.URL{
		Scheme: "https",
		Host:   "api.esa.io",
		Path:   "/",
	}
	return &ClientImpl{
		accessToken: accessToken,
		teamName:    teamName,
		options:     ops,
		baseURL:     baseURL,
	}
}

func (c *ClientImpl) httpClient() *http.Client {
	if c.options.httpClient != nil {
		return c.options.httpClient
	}
	return http.DefaultClient
}

var defaultUserAgent = fmt.Sprintf("EsaGoClient/%s (%s)", version, runtime.Version())

func (c *ClientImpl) userAgent() string {
	if c.options.userAgent != "" {
		return c.options.userAgent
	}
	return defaultUserAgent
}

// SetTeam set team
func (c *ClientImpl) SetTeam(teamName string) {
	c.teamName = teamName
}

// SetAccessToken set accessToken
func (c *ClientImpl) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}

// AppendOption append option
func (c *ClientImpl) AppendOption(options ...BaseOption) {
	ops := c.options
	for _, o := range options {
		o(&ops)
	}
	c.options = ops
}

// SetOption set option
func (c *ClientImpl) SetOption(options ...BaseOption) {
	ops := baseOptions{}
	for _, o := range options {
		o(&ops)
	}
	c.options = ops
}

// ClearOptions clear options
func (c *ClientImpl) ClearOptions() {
	c.options = baseOptions{}
}

// OverwriteBaseURL overwrite baseURL
func (c *ClientImpl) OverwriteBaseURL(u url.URL) {
	c.baseURL = u
}
