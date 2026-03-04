package groupmappings

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/federations/saml/testdata"
)

const (
	federationGroupMappingsURL        = "v1/federations/saml/123/group-mappings"
	federationExternalGroupMappingURL = "v1/federations/saml/123/group-mappings/456/external-groups/external-group"
)

func TestList(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse *GroupMappingsResponse
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+federationGroupMappingsURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGroupMappingsResponse)
						return resp, nil
					})
			},
			expectedResponse: &GroupMappingsResponse{
				GroupMappings: []GroupMapping{
					{
						InternalGroupID: "456",
						ExternalGroupID: "external-group",
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+federationGroupMappingsURL,
					func(r *http.Request) (*http.Response, error) {
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

			api := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(api.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := api.List(ctx, "123")

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name          string
		prepare       func()
		input         GroupMappingsRequest
		expectedError error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+federationGroupMappingsURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			input: GroupMappingsRequest{
				GroupMappings: []GroupMapping{
					{
						InternalGroupID: "456",
						ExternalGroupID: "external-group",
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+federationGroupMappingsURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusForbidden, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			input: GroupMappingsRequest{
				GroupMappings: []GroupMapping{
					{
						InternalGroupID: "456",
						ExternalGroupID: "external-group",
					},
				},
			},
			expectedError: iamerrors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			api := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(api.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := api.Update(ctx, "123", tt.input)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name          string
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+federationExternalGroupMappingURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+federationExternalGroupMappingURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusForbidden, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			expectedError: iamerrors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			api := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(api.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := api.Add(ctx, "123", "456", "external-group")

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name          string
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+federationExternalGroupMappingURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+federationExternalGroupMappingURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusForbidden, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			expectedError: iamerrors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			api := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(api.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := api.Delete(ctx, "123", "456", "external-group")

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name           string
		prepare        func()
		expectedExists bool
		expectedError  error
	}{
		{
			name: "exists",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodHead, testdata.TestURL+federationExternalGroupMappingURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedExists: true,
			expectedError:  nil,
		},
		{
			name: "not found",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodHead, testdata.TestURL+federationExternalGroupMappingURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNotFound, testdata.TestUserOrGroupNotFoundErr)
						return resp, nil
					})
			},
			expectedExists: false,
			expectedError:  nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodHead, testdata.TestURL+federationExternalGroupMappingURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusForbidden, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			expectedExists: false,
			expectedError:  iamerrors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			api := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(api.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			exists, err := api.Exists(ctx, "123", "456", "external-group")

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedExists, exists)
		})
	}
}
