package users

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

// Service is used to communicate with the Users API.
type Service struct {
	baseClient *client.BaseClient
}

// New initialises Service with the given client.
func New(baseClient *client.BaseClient) *Service {
	return &Service{
		baseClient: baseClient,
	}
}

// List returns a list of Users for the account.
func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	path, err := url.JoinPath(apiVersion, "users")
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

	var users ListResponse
	err = client.UnmarshalJSON(response, &users)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &users, nil
}

// Get returns an info of User with the selectel userID.
func (s *Service) Get(ctx context.Context, userID string) (*GetResponse, error) {
	if userID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "users", userID)
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

	var user GetResponse
	err = client.UnmarshalJSON(response, &user)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &user, nil
}

// Create creates a new User.
func (s *Service) Create(ctx context.Context, input CreateRequest) (*CreateResponse, error) {
	if input.Email == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserEmailRequired, Desc: "No email for User was provided."}
	}

	path, err := url.JoinPath(apiVersion, "users")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(&createRequest{
		AuthType:          input.AuthType,
		Email:             input.Email,
		Federation:        input.Federation,
		Roles:             input.Roles,
		GroupIds:          input.GroupIDs,
		SubscriptionsOnly: false,      // Issue, should be hardcoded
		Subscriptions:     []string{}, // Issue, should be hardcoded
	})
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

	var createdUser CreateResponse
	err = client.UnmarshalJSON(response, &createdUser)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &createdUser, nil
}

// Delete deletes a User from the account.
func (s *Service) Delete(ctx context.Context, userID string) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "users", userID)
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

// ResendInvite sends a confirmation email again.
func (s *Service) ResendInvite(ctx context.Context, userID string) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "users", userID, "resend_invite")
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodPatch,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}

// AssignRoles adds new roles for a User with the given userID.
func (s *Service) AssignRoles(ctx context.Context, userID string, roles []roles.Role) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	if len(roles) == 0 {
		return iamerrors.Error{Err: iamerrors.ErrUserRolesRequired, Desc: "No roles for User was provided."}
	}

	return s.manageRoles(ctx, http.MethodPut, userID, roles)
}

// UnassignRoles removes roles from a User with the given userID.
func (s *Service) UnassignRoles(ctx context.Context, userID string, roles []roles.Role) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}
	if len(roles) == 0 {
		return iamerrors.Error{Err: iamerrors.ErrUserRolesRequired, Desc: "No roles for User was provided."}
	}

	return s.manageRoles(ctx, http.MethodDelete, userID, roles)
}

func (s *Service) manageRoles(ctx context.Context, method string, userID string, roles []roles.Role) error {
	path, err := url.JoinPath(apiVersion, "users", userID, "roles")
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
