package federations

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/federations/testdata"
)

const (
	federationsURL   = "v1/federations/saml"
	federationsIDURL = "v1/federations/saml/123"
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
				[]FederationWithIDs{
					{
						ID:        "123",
						AccountID: "123",
						Federation: Federation{
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
				FederationWithIDs: FederationWithIDs{
					ID:        "123",
					AccountID: "123",
					Federation: Federation{
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
				Federation: Federation{
					Name:               "test_name",
					Description:        "test_description",
					Issuer:             "test_issuer",
					SSOUrl:             "test_sso_url",
					SignAuthnRequests:  true,
					ForceAuthn:         true,
					SessionMaxAgeHours: 1,
				},
			},
			expectedResponse: &CreateResponse{
				FederationWithIDs: FederationWithIDs{
					ID:        "123",
					AccountID: "123",
					Federation: Federation{
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
					http.MethodPost, testdata.TestURL+federationsURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusForbidden, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			input: CreateRequest{
				Federation: Federation{
					Name:               "test_name",
					Description:        "test_description",
					Issuer:             "test_issuer",
					SSOUrl:             "test_sso_url",
					SignAuthnRequests:  true,
					ForceAuthn:         true,
					SessionMaxAgeHours: 1,
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
		input            UpdateRequest
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
			input: UpdateRequest{
				Federation: Federation{
					Name:               "test_name",
					Description:        "test_description",
					Issuer:             "test_issuer",
					SSOUrl:             "test_sso_url",
					SignAuthnRequests:  true,
					ForceAuthn:         true,
					SessionMaxAgeHours: 1,
				},
			},
			expectedResponse: &GetResponse{
				FederationWithIDs: FederationWithIDs{
					ID:        "123",
					AccountID: "123",
					Federation: Federation{
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
					http.MethodPatch, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusForbidden, testdata.TestDoRequestErr)
						return resp, nil
					})
			},
			input: UpdateRequest{
				Federation: Federation{
					Name:               "test_name",
					Description:        "test_description",
					Issuer:             "test_issuer",
					SSOUrl:             "test_sso_url",
					SignAuthnRequests:  true,
					ForceAuthn:         true,
					SessionMaxAgeHours: 1,
				},
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

			ctx := context.Background()
			err := federationsAPI.Update(ctx, "123", tt.input)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		name          string
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodHead, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodHead, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
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
			err := federationsAPI.Check(ctx, "123")

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
