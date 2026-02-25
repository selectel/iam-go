package saml

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
	federationsURL                    = "v1/federations/saml"
	federationsIDURL                  = "v1/federations/saml/123"
	federationGroupMappingsURL        = "v1/federations/saml/123/group-mappings"
	federationExternalGroupMappingURL = "v1/federations/saml/123/group-mappings/456/external-groups/external-group"
)

// Convenience vars for bool values.
var (
	iTrue = true
)

func TestList(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse *ListResponse
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+federationsURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestListFederationsResponse)
						return resp, nil
					})
			},
			expectedResponse: &ListResponse{
				[]Federation{
					{
						ID:                 "123",
						AccountID:          "123",
						Name:               "test_name",
						Description:        "test_description",
						Issuer:             "test_issuer",
						SSOUrl:             "test_sso_url",
						SignAuthnRequests:  true,
						ForceAuthn:         true,
						SessionMaxAgeHours: 1,
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+federationsURL, func(r *http.Request) (*http.Response, error) {
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := federationsAPI.List(ctx)

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse *GetResponse
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetFederationResponse)
						return resp, nil
					})
			},
			expectedResponse: &GetResponse{
				Federation: Federation{
					ID:                 "123",
					AccountID:          "123",
					Name:               "test_name",
					Description:        "test_description",
					Alias:              "test_alias",
					Issuer:             "test_issuer",
					SSOUrl:             "test_sso_url",
					SignAuthnRequests:  true,
					ForceAuthn:         true,
					SessionMaxAgeHours: 1,
					AutoUsersCreation:  true,
					EnableGroupMapping: true,
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := federationsAPI.Get(ctx, "123")

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		input            CreateRequest
		expectedResponse *CreateResponse
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+federationsURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusCreated, testdata.TestCreateFederationResponse)
						return resp, nil
					})
			},
			input: CreateRequest{
				Name:               "test_name",
				Description:        "test_description",
				Issuer:             "test_issuer",
				SSOUrl:             "test_sso_url",
				SignAuthnRequests:  true,
				ForceAuthn:         true,
				SessionMaxAgeHours: 1,
			},
			expectedResponse: &CreateResponse{
				Federation: Federation{
					ID:                 "123",
					AccountID:          "123",
					Name:               "test_name",
					Description:        "test_description",
					Alias:              "test_alias",
					Issuer:             "test_issuer",
					SSOUrl:             "test_sso_url",
					SignAuthnRequests:  true,
					ForceAuthn:         true,
					SessionMaxAgeHours: 1,
					AutoUsersCreation:  true,
					EnableGroupMapping: true,
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+federationsURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusForbidden, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			input: CreateRequest{
				Name:               "test_name",
				Description:        "test_description",
				Issuer:             "test_issuer",
				SSOUrl:             "test_sso_url",
				SignAuthnRequests:  true,
				ForceAuthn:         true,
				SessionMaxAgeHours: 1,
			},
			expectedResponse: nil,
			expectedError:    iamerrors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := federationsAPI.Create(ctx, tt.input)

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse *GetResponse
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetFederationResponse)
						return resp, nil
					})
			},
			expectedResponse: &GetResponse{
				Federation: Federation{
					ID:                 "123",
					AccountID:          "123",
					Name:               "test_name",
					Description:        "test_description",
					Alias:              "test_alias",
					Issuer:             "test_issuer",
					SSOUrl:             "test_sso_url",
					SignAuthnRequests:  true,
					ForceAuthn:         true,
					SessionMaxAgeHours: 1,
					AutoUsersCreation:  true,
					EnableGroupMapping: true,
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
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
			require := require.New(t)

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			desc := "test_description"
			ctx := context.Background()
			err := federationsAPI.Update(ctx, "123", UpdateRequest{
				Name:               "test_name",
				Description:        &desc,
				Issuer:             "test_issuer",
				SSOUrl:             "test_sso_url",
				SignAuthnRequests:  &iTrue,
				ForceAuthn:         &iTrue,
				SessionMaxAgeHours: 1,
			})

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
					http.MethodDelete, testdata.TestURL+federationsIDURL,
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
					http.MethodDelete, testdata.TestURL+federationsIDURL,
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := federationsAPI.Delete(ctx, "123")

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
					http.MethodHead, testdata.TestURL+federationsIDURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, "")
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
					http.MethodHead, testdata.TestURL+federationsIDURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNotFound, testdata.TestFederationNotFoundErr)
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
					http.MethodHead, testdata.TestURL+federationsIDURL,
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			exists, err := federationsAPI.Exists(ctx, "123")

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedExists, exists)
		})
	}
}

func TestPreview(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse *FederationPreview
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+federationsIDURL+"/preview",
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestPreviewFederationResponse)
						return resp, nil
					})
			},
			expectedResponse: &FederationPreview{
				ID:          "123",
				Name:        "test_name",
				Description: "test_description",
				Alias:       "test_alias",
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+federationsIDURL+"/preview",
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := federationsAPI.Preview(ctx, "123")

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestGetGroupMappings(t *testing.T) {
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := federationsAPI.GetGroupMappings(ctx, "123")

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestUpdateGroupMappings(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		input            GroupMappingsRequest
		expectedResponse *GroupMappingsResponse
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+federationGroupMappingsURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGroupMappingsResponse)
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
			expectedResponse: nil,
			expectedError:    iamerrors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := federationsAPI.UpdateGroupMappings(ctx, "123", tt.input)

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestAddExternalGroupMapping(t *testing.T) {
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := federationsAPI.AddExternalGroupMapping(ctx, "123", "456", "external-group")

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestDeleteExternalGroupMapping(t *testing.T) {
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := federationsAPI.DeleteExternalGroupMapping(ctx, "123", "456", "external-group")

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestExternalGroupMappingExists(t *testing.T) {
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

			federationsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(federationsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			exists, err := federationsAPI.ExternalGroupMappingExists(ctx, "123", "456", "external-group")

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedExists, exists)
		})
	}
}
