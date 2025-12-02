package roles

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/roles/testdata"
)

const rolesURL = "iam/v1/roles"

func TestList(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse *ListResponse
		expectedError    error
	}{
		{
			name: "return roles list",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+rolesURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestListRolesResponse)
						return resp, nil
					})
			},
			expectedResponse: &ListResponse{Roles: []AvailableRole{
				{
					AvailableInOnboarding: true,
					Category:              "general",
					Description:           "Test role",
					ID:                    "role-id",
					Scopes:                []string{"account"},
					SubjectTypes:          []string{"user"},
					Deprecated:            false,
				},
			}},
			expectedError: nil,
		},
		{
			name: "return error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+rolesURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusForbidden, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			expectedResponse: nil,
			expectedError:    iamerrors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			rolesAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(rolesAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := rolesAPI.List(ctx)

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}
