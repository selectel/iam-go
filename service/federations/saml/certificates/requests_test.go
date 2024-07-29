package certificates

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/federations/saml/certificates/testdata"
)

const (
	certificatesURL = "v1/federations/saml/123/certificates"
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
					http.MethodGet, testdata.TestURL+certificatesURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestListCertificatesResponse)
						return resp, nil
					})
			},
			expectedResponse: &ListResponse{
				[]Certificate{
					{
						ID:           "123",
						AccountID:    "123",
						FederationID: "123",
						Name:         "test_name",
						Description:  "test_description",
						NotBefore:    "2021-01-01T00:00:00Z",
						NotAfter:     "2022-01-01T00:00:00Z",
						Fingerprint:  "test_fingerprint",
						Data:         "test_data",
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+certificatesURL, func(r *http.Request) (*http.Response, error) {
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

			certificatesAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})
			httpmock.ActivateNonDefault(certificatesAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := certificatesAPI.List(ctx, "123")

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
					http.MethodGet, testdata.TestURL+certificatesURL+"/123",
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetCertificateResponse)
						return resp, nil
					})
			},
			expectedResponse: &GetResponse{
				Certificate: Certificate{
					ID:           "123",
					AccountID:    "123",
					FederationID: "123",
					Name:         "test_name",
					Description:  "test_description",
					NotBefore:    "2021-01-01T00:00:00Z",
					NotAfter:     "2022-01-01T00:00:00Z",
					Fingerprint:  "test_fingerprint",
					Data:         "test_data",
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+certificatesURL+"/123",
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

			certificatesAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})
			httpmock.ActivateNonDefault(certificatesAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := certificatesAPI.Get(ctx, "123", "123")

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse *CreateResponse
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+certificatesURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusCreated, testdata.TestCreateCertificateResponse)
						return resp, nil
					})
			},
			expectedResponse: &CreateResponse{
				Certificate: Certificate{
					ID:           "123",
					AccountID:    "123",
					FederationID: "123",
					Name:         "test_name",
					Description:  "test_description",
					NotBefore:    "2021-01-01T00:00:00Z",
					NotAfter:     "2022-01-01T00:00:00Z",
					Fingerprint:  "test_fingerprint",
					Data:         "test_data",
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+certificatesURL, func(r *http.Request) (*http.Response, error) {
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

			certificatesAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})
			httpmock.ActivateNonDefault(certificatesAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := certificatesAPI.Create(ctx, "123", CreateRequest{
				Name:        "test_name",
				Description: "test_description",
				Data:        "test_data",
			})

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse *UpdateResponse
		expectedError    error
	}{
		{
			name: "ok",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+certificatesURL+"/123",
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetCertificateResponse)
						return resp, nil
					})
			},
			expectedResponse: &UpdateResponse{
				Certificate: Certificate{
					ID:           "123",
					AccountID:    "123",
					FederationID: "123",
					Name:         "test_name",
					Description:  "test_description",
					NotBefore:    "2021-01-01T00:00:00Z",
					NotAfter:     "2022-01-01T00:00:00Z",
					Fingerprint:  "test_fingerprint",
					Data:         "test_data",
				},
			},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+certificatesURL+"/123",
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

			certificatesAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})
			httpmock.ActivateNonDefault(certificatesAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			desc := "test_description"
			actual, err := certificatesAPI.Update(ctx, "123", "123", UpdateRequest{
				Name:        "test_name",
				Description: &desc,
			})

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
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
					http.MethodDelete, testdata.TestURL+certificatesURL+"/123",
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
					http.MethodDelete, testdata.TestURL+certificatesURL+"/123",
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

			certificatesAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})
			httpmock.ActivateNonDefault(certificatesAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := certificatesAPI.Delete(ctx, "123", "123")

			require.ErrorIs(err, tt.expectedError)
		})
	}
}
