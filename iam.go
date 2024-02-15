package iam

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/selectel/iam-go/iamerrors"
	baseclient "github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/ec2"
	"github.com/selectel/iam-go/service/serviceusers"
	"github.com/selectel/iam-go/service/users"
)

const (
	// appName represents an application name.
	appName = "iam-go"

	// defaultIAMApiURL represents a default Selectel IAM API URL.
	defaultIAMApiURL = "https://api.selectel.ru"

	// defaultHTTPTimeout represents the default timeout (in seconds) for HTTP requests.
	defaultHTTPTimeout = 120

	// defaultMaxIdleConns represents the maximum number of idle (keep-alive) connections.
	defaultMaxIdleConns = 100

	// defaultIdleConnTimeout represents the maximum amount of time an idle (keep-alive) connection will remain
	// idle before closing itself.
	defaultIdleConnTimeout = 100

	// defaultTLSHandshakeTimeout represents the default timeout (in seconds) for TLS handshake.
	defaultTLSHandshakeTimeout = 60

	// defaultExpectContinueTimeout represents the default amount of time to wait for a server's first
	// response headers.
	defaultExpectContinueTimeout = 1
)

// Client stores the configuration, which is needed to make requests to the IAM API.
type Client struct {
	// authOpts contains data to authenticate against Selectel IAM API.
	authOpts *AuthOpts

	// baseClient contains the configuration of the Client.
	baseClient *baseclient.BaseClient

	// Users instance is used to make requests against Selectel IAM API.
	Users *users.Users

	// ServiceUsers instance is used to make requests against Selectel IAM API.
	ServiceUsers *serviceusers.ServiceUsers

	// EC2 instance is used to make requests against Selectel IAM API.
	EC2 *ec2.EC2
}

type AuthOpts struct {
	KeystoneToken string
}

type Option func(*Client)

// WithAPIUrl is a functional parameter for Client, used to set IAM API URL.
func WithAPIUrl(url string) Option {
	return func(c *Client) {
		c.baseClient.APIUrl = url
	}
}

// WithCustomHTTPClient is a functional parameter for Client, used to set a custom HTTP client.
func WithCustomHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.baseClient.HTTPClient = httpClient
	}
}

// WithAuthOpts is a functional parameter for Client, used to set on of implementations of AuthType.
func WithAuthOpts(authOpts *AuthOpts) Option {
	return func(c *Client) {
		c.authOpts = authOpts
	}
}

// WithUserAgentPrefix is a functional parameter for Client, used to set a custom prefix.
func WithUserAgentPrefix(prefix string) Option {
	return func(c *Client) {
		c.baseClient.UserAgentPrefix = prefix
	}
}

// New returns a new instance of Client for the v1 IAM API.
func New(opts ...Option) (*Client, error) {
	c := &Client{baseClient: &baseclient.BaseClient{}}

	for _, opt := range opts {
		opt(c)
	}

	if !c.validateAndSetAuthMethod() {
		return nil, iamerrors.Error{Err: iamerrors.ErrClientNoAuthOpts, Desc: "No AuthOpts was passed"}
	}

	if c.baseClient.APIUrl == "" {
		c.baseClient.APIUrl = defaultIAMApiURL
	}

	if c.baseClient.HTTPClient == nil {
		c.baseClient.HTTPClient = &http.Client{
			Timeout:   defaultHTTPTimeout * time.Second,
			Transport: newHTTPTransport(),
		}
	}

	appVersion := findModuleVersion()
	userAgent := appName + "/" + appVersion
	if c.baseClient.UserAgentPrefix == "" {
		c.baseClient.UserAgent = userAgent
	} else {
		c.baseClient.UserAgent = c.baseClient.UserAgentPrefix + " " + userAgent
	}

	c.Users = users.New(c.baseClient)
	c.ServiceUsers = serviceusers.New(c.baseClient)
	c.EC2 = ec2.New(c.baseClient)

	return c, nil
}

func (c *Client) validateAndSetAuthMethod() bool {
	if c.authOpts == nil {
		return false
	}
	if c.authOpts.KeystoneToken != "" {
		c.baseClient.AuthMethod = &baseclient.KeystoneTokenAuth{
			KeystoneToken: c.authOpts.KeystoneToken,
		}
		return true
	}
	return false
}

func newHTTPTransport() *http.Transport {
	return &http.Transport{
		MaxIdleConns:          defaultMaxIdleConns,
		IdleConnTimeout:       defaultIdleConnTimeout * time.Second,
		TLSHandshakeTimeout:   defaultTLSHandshakeTimeout * time.Second,
		ExpectContinueTimeout: defaultExpectContinueTimeout * time.Second,
	}
}

func findModuleVersion() string {
	moduleName := "github.com/selectel/" + appName

	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range info.Deps {
			if dep.Path == moduleName {
				return dep.Version
			}
		}
	}
	return "v0.1.0"
}
