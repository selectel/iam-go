package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/users/testdata"
)

const (
	usersURL        = "iam/v1/users"
	usersIDURL      = "iam/v1/users/123"
	rolesURL        = "iam/v1/users/123/roles"
	resendInviteURL = "iam/v1/users/123/resend_invite"
)

func TestList(t *testing.T) {
	tests := []struct {
		name             string
		prepare          func()
		expectedResponse []User
		expectedError    error
	}{
		{
			name: "Test List return output",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+usersURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetUsersResponse)
						return resp, nil
					})
			},
			expectedResponse: []User{
				{
					AuthType:   "local",
					KeystoneID: "123",
					ID:         "123",
					Roles: []Role{
						{Scope: Account, RoleName: Member},
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "Test List return error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+usersURL, func(r *http.Request) (*http.Response, error) {
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

			usersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(usersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := usersAPI.List(ctx)

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse *User
		expectedError    error
	}{
		{
			name: "Test Get return output",
			args: args{
				userID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+usersIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetUserResponse)
						return resp, nil
					})
			},
			expectedResponse: &User{
				AuthType:   "local",
				KeystoneID: "123",
				ID:         "123",
				Roles: []Role{
					{Scope: Account, RoleName: Member},
				},
			},
			expectedError: nil,
		},
		{
			name: "Test Get return error",
			args: args{
				userID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+usersIDURL, func(r *http.Request) (*http.Response, error) {
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

			usersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(usersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := usersAPI.Get(ctx, tt.args.userID)

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		userID string
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
				userID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+usersIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "Test Delete return error",
			args: args{
				userID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+usersIDURL, func(r *http.Request) (*http.Response, error) {
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

			usersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(usersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := usersAPI.Delete(ctx, tt.args.userID)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		authType   AuthType
		email      string
		federation Federation
		roles      []Role
	}
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse *User
		expectedError    error
	}{
		{
			name: "Test Create return output",
			args: args{
				authType: "federated",
				email:    "test@mail",
				federation: Federation{
					ExternalID: "123",
					ID:         "123",
				},
				roles: []Role{
					{Scope: Account, RoleName: Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+usersURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestCreateUserResponse)
						return resp, nil
					})
			},
			expectedResponse: &User{
				AuthType: "federated",
				Federation: &Federation{
					ExternalID: "123",
					ID:         "123",
				},
				ID:         "123",
				KeystoneID: "123",
				Roles: []Role{
					{Scope: Account, RoleName: Member},
				},
			},
			expectedError: nil,
		},
		{
			name: "Test Create return error",
			args: args{
				authType: "federated",
				email:    "test@mail",
				federation: Federation{
					ExternalID: "123",
					ID:         "123",
				},
				roles: []Role{
					{Scope: Account, RoleName: Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+usersURL, func(r *http.Request) (*http.Response, error) {
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

			usersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(usersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()
			ctx := context.Background()

			//nolint:gosec // This is just a test
			actual, err := usersAPI.Create(ctx, CreateRequest{
				AuthType:   tt.args.authType,
				Email:      tt.args.email,
				Federation: &tt.args.federation,
				Roles:      tt.args.roles,
			})
			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestResendInvite(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "Test ResendInvite return nil",
			args: args{
				userID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+resendInviteURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "Test ResendInvite return error",
			args: args{
				userID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+resendInviteURL, func(r *http.Request) (*http.Response, error) {
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

			usersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(usersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := usersAPI.ResendInvite(ctx, tt.args.userID)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestAssignRoles(t *testing.T) {
	type args struct {
		userID string
		roles  []Role
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "Test AssignRoles return output",
			args: args{
				userID: "123",
				roles: []Role{
					{Scope: Account, RoleName: Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+rolesURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "Test AssignRoles return error",
			args: args{
				userID: "123",
				roles: []Role{
					{Scope: Account, RoleName: Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+rolesURL, func(r *http.Request) (*http.Response, error) {
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

			usersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(usersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := usersAPI.AssignRoles(ctx, tt.args.userID, tt.args.roles)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestUnassignRoles(t *testing.T) {
	type args struct {
		userID string
		roles  []Role
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "Test UnassignRoles return output",
			args: args{
				userID: "123",
				roles: []Role{
					{Scope: Account, RoleName: Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+rolesURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "Test UnassignRoles return error",
			args: args{
				userID: "123",
				roles: []Role{
					{Scope: Account, RoleName: Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+rolesURL, func(r *http.Request) (*http.Response, error) {
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

			usersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(usersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := usersAPI.UnassignRoles(ctx, tt.args.userID, tt.args.roles)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}
