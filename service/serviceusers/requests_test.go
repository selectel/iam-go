package serviceusers

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/serviceusers/testdata"
)

const (
	serviceUsersURL      = "iam/v1/service_users"
	serviceUsersIDURL    = "iam/v1/service_users/123"
	serviceUsersRolesURL = "iam/v1/service_users/123/roles"
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
					http.MethodGet, testdata.TestURL+serviceUsersURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestListUsersResponse)
						return resp, nil
					})
			},
			expectedResponse: &ListResponse{[]ServiceUser{
				{
					Name:    "test",
					Enabled: true,
					ID:      "123",
					Roles: []roles.Role{
						{Scope: roles.Account, RoleName: roles.Member},
					},
				},
			}},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+serviceUsersURL, func(r *http.Request) (*http.Response, error) {
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

			serviceUsersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(serviceUsersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := serviceUsersAPI.List(ctx)

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
		expectedResponse *GetResponse
		expectedError    error
	}{
		{
			name: "Test Get return output",
			args: args{
				userID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+serviceUsersIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetUserResponse)
						return resp, nil
					})
			},
			expectedResponse: &GetResponse{
				ServiceUser: ServiceUser{
					Name:    "test",
					Enabled: true,
					ID:      "123",
					Roles: []roles.Role{
						{Scope: roles.Account, RoleName: roles.Member},
					},
				},
				Groups: []Group{
					{Name: "123", ID: "96a60e7b9e9e48308eed46269f9a147b", Roles: []roles.Role{}},
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
					http.MethodGet, testdata.TestURL+serviceUsersIDURL, func(r *http.Request) (*http.Response, error) {
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

			serviceUsersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(serviceUsersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := serviceUsersAPI.Get(ctx, tt.args.userID)

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
					http.MethodDelete, testdata.TestURL+serviceUsersIDURL,
					func(r *http.Request) (*http.Response, error) {
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
					http.MethodDelete, testdata.TestURL+serviceUsersIDURL,
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

			serviceUsersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(serviceUsersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := serviceUsersAPI.Delete(ctx, tt.args.userID)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

//nolint:funlen // This is a test function.
func TestCreate(t *testing.T) {
	type args struct {
		enabled  bool
		name     string
		password string
		roles    []roles.Role
	}
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse *CreateResponse
		expectedError    error
	}{
		{
			name: "Test Create return output",
			args: args{
				enabled:  true,
				name:     "test",
				password: "Qazwsxedc123",
				roles: []roles.Role{
					{Scope: roles.Account, RoleName: roles.Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+serviceUsersURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestCreateUserResponse)
						return resp, nil
					})
			},
			expectedResponse: &CreateResponse{
				ServiceUser: ServiceUser{
					Name:    "test",
					Enabled: true,
					ID:      "123",
					Roles: []roles.Role{
						{Scope: roles.Account, RoleName: roles.Member},
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "Test Create return insecure password",
			args: args{
				enabled:  true,
				name:     "test",
				password: "123",
				roles: []roles.Role{
					{Scope: roles.Account, RoleName: roles.Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+serviceUsersURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(
							http.StatusBadRequest,
							testdata.TestCreateUserInsecurePasswordErr,
						)
						return resp, nil
					})
			},
			expectedResponse: nil,
			expectedError:    iamerrors.ErrRequestValidationError,
		},
		{
			name: "Test Create return error",
			args: args{
				enabled:  true,
				name:     "test",
				password: "123",
				roles: []roles.Role{
					{Scope: roles.Account, RoleName: roles.Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+serviceUsersURL, func(r *http.Request) (*http.Response, error) {
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

			serviceUsersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(serviceUsersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := serviceUsersAPI.Create(ctx, CreateRequest{
				Name:     tt.args.name,
				Enabled:  tt.args.enabled,
				Password: tt.args.password,
				Roles:    tt.args.roles,
			})

			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		userID   string
		enabled  bool
		name     string
		password string
	}
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse *UpdateResponse
		expectedError    error
	}{
		{
			name: "Test Update return output",
			args: args{
				userID:  "123",
				enabled: true,
				name:    "test1",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+serviceUsersIDURL,
					func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestUpdateUserResponse)
						return resp, nil
					})
			},
			expectedResponse: &UpdateResponse{
				ServiceUser: ServiceUser{
					Name:    "test1",
					Enabled: true,
					ID:      "123",
				},
			},
			expectedError: nil,
		},
		{
			name: "Test Update return error",
			args: args{
				userID:   "123",
				enabled:  true,
				name:     "test",
				password: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+serviceUsersIDURL,
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

			serviceUsersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(serviceUsersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()
			tt.prepare()

			ctx := context.Background()
			actual, err := serviceUsersAPI.Update(ctx, tt.args.userID, UpdateRequest{
				Enabled:  tt.args.enabled,
				Name:     tt.args.name,
				Password: tt.args.password,
			})

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestAssignRoles(t *testing.T) {
	type args struct {
		userID string
		roles  []roles.Role
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
				roles: []roles.Role{
					{Scope: roles.Account, RoleName: roles.Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+serviceUsersRolesURL,
					func(r *http.Request) (*http.Response, error) {
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
				roles: []roles.Role{
					{Scope: roles.Account, RoleName: roles.Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+serviceUsersRolesURL,
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

			serviceUsersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(serviceUsersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := serviceUsersAPI.AssignRoles(ctx, tt.args.userID, tt.args.roles)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestUnassignRoles(t *testing.T) {
	type args struct {
		userID string
		roles  []roles.Role
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
				roles: []roles.Role{
					{Scope: roles.Account, RoleName: roles.Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+serviceUsersRolesURL,
					func(r *http.Request) (*http.Response, error) {
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
				roles: []roles.Role{
					{Scope: roles.Account, RoleName: roles.Member},
				},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+serviceUsersRolesURL,
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

			serviceUsersAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(serviceUsersAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()
			tt.prepare()

			ctx := context.Background()
			err := serviceUsersAPI.UnassignRoles(ctx, tt.args.userID, tt.args.roles)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}
