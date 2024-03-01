package serviceusers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/roles"
)

const apiVersion = "iam/v1"

// ServiceUsers is used to communicate with the Service Users API.
type ServiceUsers struct {
	baseClient *client.BaseClient
}

// Initialises ServiceUsers with the given client.
func New(baseClient *client.BaseClient) *ServiceUsers {
	return &ServiceUsers{
		baseClient: baseClient,
	}
}

// List returns a list of Service Users for the account.
func (su *ServiceUsers) List(ctx context.Context) ([]ServiceUser, error) {
	path, err := url.JoinPath(apiVersion, "service_users")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := su.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var users listResponse
	err = client.UnmarshalJSON(response, &users)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return users.Users, nil
}

// Get returns an info about Service User with the selectel userID.
func (su *ServiceUsers) Get(ctx context.Context, userID string) (*ServiceUser, error) {
	if userID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "service_users", userID)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := su.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var user ServiceUser
	err = client.UnmarshalJSON(response, &user)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &user, nil
}

// Create creates a new Service User.
func (su *ServiceUsers) Create(ctx context.Context, input CreateRequest) (*ServiceUser, error) {
	if input.Name == "" {
		return nil, iamerrors.Error{
			Err: iamerrors.ErrServiceUserNameRequired, Desc: "No name for Service User was provided.",
		}
	}
	if input.Password == "" {
		return nil, iamerrors.Error{
			Err: iamerrors.ErrServiceUserPasswordRequired, Desc: "No password for Service User was provided.",
		}
	}

	path, err := url.JoinPath(apiVersion, "service_users")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	body, err := json.Marshal(&createRequest{
		Enabled:  input.Enabled,
		Name:     input.Name,
		Password: input.Password,
		Roles:    input.Roles,
	})
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := su.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPost,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var createdUser ServiceUser
	err = client.UnmarshalJSON(response, &createdUser)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &createdUser, nil
}

// Delete deletes a Service User from the account.
func (su *ServiceUsers) Delete(ctx context.Context, userID string) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "service_users", userID)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = su.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodDelete,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}

// Update updates the info for a Service User with the given userID.
func (su *ServiceUsers) Update(ctx context.Context, userID string, input UpdateRequest) (*ServiceUser, error) {
	if userID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "service_users", userID)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(&updateRequest{
		Enabled:  &input.Enabled,
		Name:     input.Name,
		Password: input.Password,
	})
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := su.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPatch,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var updatedUser ServiceUser
	err = client.UnmarshalJSON(response, &updatedUser)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &updatedUser, nil
}

// AssignRoles adds new roles for a Service User with the given userID.
func (su *ServiceUsers) AssignRoles(ctx context.Context, userID string, roles []roles.Role) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}
	if len(roles) == 0 {
		return iamerrors.Error{
			Err:  iamerrors.ErrServiceUserRolesRequired,
			Desc: "No roles for Service User was provided.",
		}
	}

	return su.manageRoles(ctx, http.MethodPut, userID, roles)
}

// UnassignRoles removes roles from a Service User with the given userID.
func (su *ServiceUsers) UnassignRoles(ctx context.Context, userID string, roles []roles.Role) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}
	if len(roles) == 0 {
		return iamerrors.Error{
			Err:  iamerrors.ErrServiceUserRolesRequired,
			Desc: "No roles for Service User was provided.",
		}
	}

	return su.manageRoles(ctx, http.MethodDelete, userID, roles)
}

func (su *ServiceUsers) manageRoles(ctx context.Context, method string, userID string, roles []roles.Role) error {
	path, err := url.JoinPath(apiVersion, "service_users", userID, "roles")
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	body, err := json.Marshal(manageRolesRequest{
		Roles: roles,
	})
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = su.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: method,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}
