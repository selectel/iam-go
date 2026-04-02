package oidc

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/federations/oidc/testdata"
)

const (
	federationsURL   = "v1/federations/oidc"
	federationsIDURL = "v1/federations/oidc/123"
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
						ClientID:           "test_client_id",
						AuthURL:            "https://idp.example.com/authorize",
						TokenURL:           "https://idp.example.com/token",
						JWKSURL:            "https://idp.example.com/.well-known/jwks.json",
						SessionMaxAgeHours: 24,
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
					ClientID:           "test_client_id",
					ClientSecret:       "test_client_secret",
					AuthURL:            "https://idp.example.com/authorize",
					TokenURL:           "https://idp.example.com/token",
					JWKSURL:            "https://idp.example.com/.well-known/jwks.json",
					SessionMaxAgeHours: 24,
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
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestCreateFederationResponse)
						return resp, nil
					})
			},
			input: CreateRequest{
				Name:               "test_name",
				Description:        "test_description",
				ClientID:           "test_client_id",
				ClientSecret:       "test_client_secret",
				AuthURL:            "https://idp.example.com/authorize",
				TokenURL:           "https://idp.example.com/token",
				JWKSURL:            "https://idp.example.com/.well-known/jwks.json",
				SessionMaxAgeHours: 24,
			},
			expectedResponse: &CreateResponse{
				Federation: Federation{
					ID:                 "123",
					AccountID:          "123",
					Name:               "test_name",
					Description:        "test_description",
					Alias:              "test_alias",
					ClientID:           "test_client_id",
					ClientSecret:       "test_client_secret",
					AuthURL:            "https://idp.example.com/authorize",
					TokenURL:           "https://idp.example.com/token",
					JWKSURL:            "https://idp.example.com/.well-known/jwks.json",
					SessionMaxAgeHours: 24,
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
				ClientID:           "test_client_id",
				AuthURL:            "https://idp.example.com/authorize",
				TokenURL:           "https://idp.example.com/token",
				JWKSURL:            "https://idp.example.com/.well-known/jwks.json",
				SessionMaxAgeHours: 24,
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
		name          string
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
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
					http.MethodPatch, testdata.TestURL+federationsIDURL, func(r *http.Request) (*http.Response, error) {
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

			desc := "test_description"
			ctx := context.Background()
			err := federationsAPI.Update(ctx, "123", UpdateRequest{
				Name:        "test_name",
				Description: &desc,
				ClientID:    "test_client_id",
				AuthURL:     "https://idp.example.com/authorize",
				TokenURL:    "https://idp.example.com/token",
				JWKSURL:     "https://idp.example.com/.well-known/jwks.json",
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
