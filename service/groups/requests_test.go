package groups

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/groups/testdata"
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/users"
)

const (
	// Account scope.
	AccountScope string = "account"

	// Account/Project member.
	Member string = "member"

	groupsURL   = "iam/v1/groups"
	groupsIDURL = "iam/v1/groups/123"
	rolesURL    = "iam/v1/groups/123/roles"
	usersURL    = "iam/v1/groups/123/users"
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
					http.MethodGet, testdata.TestURL+groupsURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestListGroupsResponse)
						return resp, nil
					})
			},
			expectedResponse: &ListResponse{[]Group{
				{
					ID:          "123",
					Name:        "test_name",
					Description: "test_description",
					Roles: []roles.Role{
						{Scope: AccountScope, RoleName: Member},
					},
				},
			}},
			expectedError: nil,
		},
		{
			name: "error",
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+groupsURL, func(r *http.Request) (*http.Response, error) {
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := groupsAPI.List(ctx)

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		groupID string
	}
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse *GetResponse
		expectedError    error
	}{
		{
			name: "ok",
			args: args{
				groupID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+groupsIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestGetGroupResponse)
						return resp, nil
					})
			},
			expectedResponse: &GetResponse{
				Group: Group{
					ID:          "123",
					Name:        "test_name",
					Description: "test_description",
					Roles: []roles.Role{
						{Scope: AccountScope, RoleName: Member},
					},
				},
				ServiceUsers: []ServiceUser{
					{
						Name:    "1234",
						Enabled: false,
						ID:      "c1f50a57fc95438aafe1e2fe87a781c2",
					},
				},
				Users: []User{{
					AuthType: users.Federated,
					Federation: &users.Federation{
						ID:         "674870b5-6ad1-478e-8384-2527507fd85d",
						ExternalID: "asdfasdf",
					},
					ID:         "999000_33333",
					KeystoneID: "2b573185f23a40c88cbb59ed74dba683",
				}},
			},
			expectedError: nil,
		},
		{
			name: "error",
			args: args{
				groupID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodGet, testdata.TestURL+groupsIDURL, func(r *http.Request) (*http.Response, error) {
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			actual, err := groupsAPI.Get(ctx, tt.args.groupID)

			require.ErrorIs(err, tt.expectedError)

			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		groupID string
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			args: args{
				groupID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+groupsIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "error",
			args: args{
				groupID: "123",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+groupsIDURL, func(r *http.Request) (*http.Response, error) {
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := groupsAPI.Delete(ctx, tt.args.groupID)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		name        string
		description string
	}
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse *CreateResponse
		expectedError    error
	}{
		{
			name: "ok",
			args: args{
				name:        "test_name",
				description: "test_description",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+groupsURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestCreateGroupResponse)
						return resp, nil
					})
			},
			expectedResponse: &CreateResponse{
				Group: Group{
					ID:          "123",
					Name:        "test_name",
					Description: "test_description",
					Roles:       []roles.Role{},
				},
				ServiceUsers: []ServiceUser{},
				Users:        []User{},
			},
			expectedError: nil,
		},
		{
			name: "error",
			args: args{
				name:        "test_name",
				description: "test_description",
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPost, testdata.TestURL+groupsURL, func(r *http.Request) (*http.Response, error) {
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()
			ctx := context.Background()

			actual, err := groupsAPI.Create(ctx, CreateRequest{
				Name:        tt.args.name,
				Description: tt.args.description,
			})
			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		groupID     string
		name        string
		description *string
	}
	testDescription := "test_description"
	tests := []struct {
		name             string
		args             args
		prepare          func()
		expectedResponse *UpdateResponse
		expectedError    error
	}{
		{
			name: "ok",
			args: args{
				groupID:     "123",
				name:        "test_name",
				description: &testDescription,
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+groupsIDURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusOK, testdata.TestCreateGroupResponse)
						return resp, nil
					})
			},
			expectedResponse: &UpdateResponse{
				Group: Group{
					ID:          "123",
					Name:        "test_name",
					Description: "test_description",
					Roles:       []roles.Role{},
				},
				ServiceUsers: []ServiceUser{},
				Users:        []User{},
			},
			expectedError: nil,
		},
		{
			name: "error",
			args: args{
				groupID:     "123",
				name:        "test_name",
				description: &testDescription,
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPatch, testdata.TestURL+groupsIDURL, func(r *http.Request) (*http.Response, error) {
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()
			ctx := context.Background()

			actual, err := groupsAPI.Update(ctx, tt.args.groupID, UpdateRequest{
				Name:        tt.args.name,
				Description: tt.args.description,
			})
			require.ErrorIs(err, tt.expectedError)
			assert.Equal(tt.expectedResponse, actual)
		})
	}
}

func TestAssignRoles(t *testing.T) {
	type args struct {
		groupID string
		roles   []roles.Role
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			args: args{
				groupID: "123",
				roles: []roles.Role{
					{Scope: AccountScope, RoleName: Member},
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
			name: "error",
			args: args{
				groupID: "123",
				roles: []roles.Role{
					{Scope: AccountScope, RoleName: Member},
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := groupsAPI.AssignRoles(ctx, tt.args.groupID, tt.args.roles)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestUnassignRoles(t *testing.T) {
	type args struct {
		groupID string
		roles   []roles.Role
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			args: args{
				groupID: "123",
				roles: []roles.Role{
					{Scope: AccountScope, RoleName: Member},
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
			name: "error",
			args: args{
				groupID: "123",
				roles: []roles.Role{
					{Scope: AccountScope, RoleName: Member},
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := groupsAPI.UnassignRoles(ctx, tt.args.groupID, tt.args.roles)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestAddUsers(t *testing.T) {
	type args struct {
		groupID  string
		usersIDs []string
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			args: args{
				groupID:  "123",
				usersIDs: []string{"user123"},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+usersURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "error",
			args: args{
				groupID:  "123",
				usersIDs: []string{"user123"},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodPut, testdata.TestURL+usersURL, func(r *http.Request) (*http.Response, error) {
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := groupsAPI.AddUsers(ctx, tt.args.groupID, tt.args.usersIDs)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}

func TestDeleteUsers(t *testing.T) {
	type args struct {
		groupID  string
		usersIDs []string
	}
	tests := []struct {
		name          string
		args          args
		prepare       func()
		expectedError error
	}{
		{
			name: "ok",
			args: args{
				groupID:  "123",
				usersIDs: []string{"user123"},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+usersURL, func(r *http.Request) (*http.Response, error) {
						resp := httpmock.NewStringResponse(http.StatusNoContent, "")
						return resp, nil
					})
			},
			expectedError: nil,
		},
		{
			name: "error",
			args: args{
				groupID:  "123",
				usersIDs: []string{"user123"},
			},
			prepare: func() {
				httpmock.RegisterResponder(
					http.MethodDelete, testdata.TestURL+usersURL, func(r *http.Request) (*http.Response, error) {
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

			groupsAPI := New(&client.BaseClient{
				HTTPClient: &http.Client{},
				APIUrl:     testdata.TestURL,
				AuthMethod: &client.KeystoneTokenAuth{KeystoneToken: testdata.TestToken},
			})

			httpmock.ActivateNonDefault(groupsAPI.baseClient.HTTPClient)
			defer httpmock.DeactivateAndReset()

			tt.prepare()

			ctx := context.Background()
			err := groupsAPI.DeleteUsers(ctx, tt.args.groupID, tt.args.usersIDs)

			require.ErrorIs(err, tt.expectedError)
		})
	}
}
