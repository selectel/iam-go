package iam

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	baseclient "github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/ec2"
	"github.com/selectel/iam-go/service/serviceusers"
	"github.com/selectel/iam-go/service/users"
)

const (
	testToken = "test-token"
	testURL   = "http://example.org/"
)

//nolint:funlen // This is a test function.
func TestNew(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name               string
		args               args
		expectedClient     func() *Client
		expectedBaseClient *baseclient.BaseClient
		expectedError      error
	}{
		{
			name: "Test NewIAMClientV1 only with TokenAuth and APIUrl",
			args: args{
				opts: []Option{
					WithAPIUrl(testURL),
					WithAuthOpts(&AuthOpts{
						KeystoneToken: testToken,
					}),
				},
			},
			expectedClient: func() *Client {
				baseClient := &baseclient.BaseClient{
					HTTPClient: &http.Client{
						Timeout:   defaultHTTPTimeout * time.Second,
						Transport: newHTTPTransport(),
					},
					APIUrl:     testURL,
					AuthMethod: &baseclient.KeystoneTokenAuth{KeystoneToken: testToken},
					UserAgent:  appName + "/" + findModuleVersion(),
				}
				return &Client{
					authOpts: &AuthOpts{
						KeystoneToken: testToken,
					},
					baseClient:   baseClient,
					Users:        users.New(baseClient),
					ServiceUsers: serviceusers.New(baseClient),
					EC2:          ec2.New(baseClient),
				}
			},
			expectedError: nil,
		},
		{
			name: "Test NewIAMClientV1 only with APIUrl",
			args: args{
				opts: []Option{
					WithAPIUrl(testURL),
				},
			},
			expectedClient: nil,
			expectedError:  iamerrors.ErrClientNoAuthOpts,
		},
		{
			name: "Test NewIAMClientV1 only with TokenAuth",
			args: args{
				opts: []Option{
					WithAuthOpts(&AuthOpts{
						KeystoneToken: testToken,
					}),
				},
			},
			expectedClient: func() *Client {
				baseClient := &baseclient.BaseClient{
					HTTPClient: &http.Client{
						Timeout:   defaultHTTPTimeout * time.Second,
						Transport: newHTTPTransport(),
					},
					APIUrl:     defaultIAMApiURL,
					AuthMethod: &baseclient.KeystoneTokenAuth{KeystoneToken: testToken},
					UserAgent:  appName + "/" + findModuleVersion(),
				}
				return &Client{
					authOpts: &AuthOpts{
						KeystoneToken: testToken,
					},
					baseClient:   baseClient,
					Users:        users.New(baseClient),
					ServiceUsers: serviceusers.New(baseClient),
					EC2:          ec2.New(baseClient),
				}
			},
			expectedError: nil,
		},
		{
			name: "Test NewIAMClientV1 only with TokenAuth and APIUrl and HTTPClient",
			args: args{
				opts: []Option{
					WithAPIUrl(testURL),
					WithAuthOpts(&AuthOpts{
						KeystoneToken: testToken,
					}),
					WithCustomHTTPClient(&http.Client{
						Timeout: 10 * time.Second,
					}),
				},
			},
			expectedClient: func() *Client {
				baseClient := &baseclient.BaseClient{
					HTTPClient: &http.Client{
						Timeout: 10 * time.Second,
					},
					APIUrl:     testURL,
					AuthMethod: &baseclient.KeystoneTokenAuth{KeystoneToken: testToken},
					UserAgent:  appName + "/" + findModuleVersion(),
				}
				return &Client{
					authOpts: &AuthOpts{
						KeystoneToken: testToken,
					},
					baseClient:   baseClient,
					Users:        users.New(baseClient),
					ServiceUsers: serviceusers.New(baseClient),
					EC2:          ec2.New(baseClient),
				}
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)
			var expected *Client
			if tt.expectedClient != nil {
				expected = tt.expectedClient()
			}
			actual, err := New(tt.args.opts...)
			require.ErrorIs(err, tt.expectedError)
			assert.Equal(expected, actual)
		})
	}
}
