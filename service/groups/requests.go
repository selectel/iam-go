package groups

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

// Service is used to communicate with the Groups API.
type Service struct {
	baseClient *client.BaseClient
}

// New Initialises Service with the given client.
func New(baseClient *client.BaseClient) *Service {
	return &Service{
		baseClient: baseClient,
	}
}

// List returns a list of Groups for the account.
func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	path, err := url.JoinPath(apiVersion, "groups")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var groups ListResponse
	err = client.UnmarshalJSON(response, &groups)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &groups, nil
}

// Get returns an info of Group with groupID.
func (s *Service) Get(ctx context.Context, groupID string) (*GetResponse, error) {
	if groupID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrGroupIDRequired, Desc: "No groupID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "groups", groupID)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var group GetResponse
	err = client.UnmarshalJSON(response, &group)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &group, nil
}

// Create creates a new Group.
func (s *Service) Create(ctx context.Context, input CreateRequest) (*CreateResponse, error) {
	if input.Name == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrGroupNameRequired, Desc: "No Name for Group was provided."}
	}

	path, err := url.JoinPath(apiVersion, "groups")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(input)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPost,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var group CreateResponse
	err = client.UnmarshalJSON(response, &group)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &group, nil
}

// Update updates exists Group.
func (s *Service) Update(ctx context.Context, groupID string, input UpdateRequest) (*UpdateResponse, error) {
	if groupID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrGroupIDRequired, Desc: "No groupID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "groups", groupID)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(input)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPatch,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var group UpdateResponse
	err = client.UnmarshalJSON(response, &group)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &group, nil
}

// Delete deletes a Group from the account.
func (s *Service) Delete(ctx context.Context, groupID string) error {
	if groupID == "" {
		return iamerrors.Error{Err: iamerrors.ErrGroupIDRequired, Desc: "No groupID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "groups", groupID)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
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

// AssignRoles adds new roles for a Group with the given groupID.
func (s *Service) AssignRoles(ctx context.Context, groupID string, roles []roles.Role) error {
	if groupID == "" {
		return iamerrors.Error{Err: iamerrors.ErrGroupIDRequired, Desc: "No groupID was provided."}
	}

	if len(roles) == 0 {
		return iamerrors.Error{Err: iamerrors.ErrGroupRolesRequired, Desc: "No roles for Group was provided."}
	}

	return s.manageRoles(ctx, http.MethodPut, groupID, roles)
}

// UnassignRoles removes roles from a Group with the given groupID.
func (s *Service) UnassignRoles(ctx context.Context, groupID string, roles []roles.Role) error {
	if groupID == "" {
		return iamerrors.Error{Err: iamerrors.ErrGroupIDRequired, Desc: "No groupID was provided."}
	}
	if len(roles) == 0 {
		return iamerrors.Error{Err: iamerrors.ErrGroupRolesRequired, Desc: "No roles for Group was provided."}
	}

	return s.manageRoles(ctx, http.MethodDelete, groupID, roles)
}

func (s *Service) manageRoles(ctx context.Context, method string, groupID string, roles []roles.Role) error {
	path, err := url.JoinPath(apiVersion, "groups", groupID, "roles")
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	body, err := json.Marshal(manageRolesRequest{
		Roles: roles,
	})
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
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

// AddUsers adds new users to a Group with the given groupID.
func (s *Service) AddUsers(ctx context.Context, groupID string, usersKeystoneIDs []string) error {
	if groupID == "" {
		return iamerrors.Error{Err: iamerrors.ErrGroupIDRequired, Desc: "No groupID was provided."}
	}

	if len(usersKeystoneIDs) == 0 {
		return iamerrors.Error{Err: iamerrors.ErrGroupUserIDsRequired, Desc: "No users for Group was provided."}
	}

	return s.manageUsers(ctx, http.MethodPut, groupID, usersKeystoneIDs)
}

// DeleteUsers removes users from a Group with the given groupID.
func (s *Service) DeleteUsers(ctx context.Context, groupID string, usersKeystoneIDs []string) error {
	if groupID == "" {
		return iamerrors.Error{Err: iamerrors.ErrGroupIDRequired, Desc: "No groupID was provided."}
	}
	if len(usersKeystoneIDs) == 0 {
		return iamerrors.Error{Err: iamerrors.ErrGroupUserIDsRequired, Desc: "No users for Group was provided."}
	}

	return s.manageUsers(ctx, http.MethodDelete, groupID, usersKeystoneIDs)
}

func (s *Service) manageUsers(ctx context.Context, method string, groupID string, usersKeystoneIDs []string) error {
	path, err := url.JoinPath(apiVersion, "groups", groupID, "users")
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	body, err := json.Marshal(manageUsersRequest{
		KeystoneIds: usersKeystoneIDs,
	})
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
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
