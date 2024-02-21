package ec2

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/ec2/testdata"
)

const (
	//nolint:gosec
	credentialsURL = "iam/v1/service_users/1/credentials"
)

func TestList(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse []Credential
		expectedError    error
	}{
		{
			name: "Test List return output",
			args: args{
				userID: "1",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+credentialsURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetCredentialsResponse)
						return resp, nil
					})
			},
			expectedResponse: []Credential{
				{
					Name:      "12345",
					ProjectID: "test-project",
					AccessKey: "test-access-key",
				},
			},
			expectedError: nil,
		},
		{
			name: "Test List return error",
			args: args{
				userID: "1",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+credentialsURL, func(r *http.Request) (*http.Response, error) {
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

			ec2CredAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(ec2CredAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actualResponse, err := ec2CredAPI.List(ctx, tt.args.userID)

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actualResponse)
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		userID    string
		name      string
		projectID string
	}
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse *CreatedCredential
		expectedError    error
	}{
		{
			name: "Test Create return output",
			args: args{
				userID:    "1",
				name:      "12345",
				projectID: "test-project",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+credentialsURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestCreateCredentialResponse)
						return resp, nil
					})
			},
			expectedResponse: &CreatedCredential{
				Name:      "12345",
				ProjectID: "test-project",
				AccessKey: "test-access-key",
				SecretKey: "test-secret-key",
			},
			expectedError: nil,
		},
		{
			name: "Test Create return error",
			args: args{
				userID:    "1",
				name:      "12345",
				projectID: "test-project",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+credentialsURL, func(r *http.Request) (*http.Response, error) {
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

			ec2CredAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(ec2CredAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actualResponse, err := ec2CredAPI.Create(ctx, tt.args.userID, tt.args.name, tt.args.projectID)

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actualResponse)
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		userID    string
		accessKey string
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "Test Delete return output",
			args: args{
				userID:    "1",
				accessKey: "test-access-key",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+credentialsURL+"/test-access-key",
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						defer resp.Body.Close()
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "Test Delete return error",
			args: args{
				userID:    "1",
				accessKey: "test-access-key",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+credentialsURL+"/test-access-key",
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

			ec2CredAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(ec2CredAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := ec2CredAPI.Delete(ctx, tt.args.userID, tt.args.accessKey)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}
