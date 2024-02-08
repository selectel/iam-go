package users

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
)

// Users is used to communicate with the Users API.
type Users struct {
	baseClient *client.BaseClient
}

// Initialises Users with the given client.
func New(baseClient *client.BaseClient) *Users {
	return &Users{
		baseClient: baseClient,
	}
}

// List returns a list of Users for the account.
func (u *Users) List(ctx context.Context) ([]User, error) {
	url, err := url.JoinPath(u.baseClient.APIUrl, "users")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := u.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		URL:    url,
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

// Get returns an info of User with the selectel userID.
func (u *Users) Get(ctx context.Context, userID string) (*User, error) {
	if userID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	url, err := url.JoinPath(u.baseClient.APIUrl, "users", userID)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := u.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		URL:    url,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var user User
	err = client.UnmarshalJSON(response, &user)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &user, nil
}

// Create creates a new User.
func (u *Users) Create(ctx context.Context, input CreateRequest) (*User, error) {
	if input.Email == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserEmailRequired, Desc: "No email for User was provided."}
	}

	url, err := url.JoinPath(u.baseClient.APIUrl, "users")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(&createRequest{
		AuthType:          input.AuthType,
		Email:             input.Email,
		Federation:        input.Federation,
		Roles:             input.Roles,
		SubscriptionsOnly: false,      // Issue, should be hardcoded
		Subscriptions:     []string{}, // Issue, should be hardcoded
	})
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := u.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPost,
		URL:    url,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var createdUser User
	err = client.UnmarshalJSON(response, &createdUser)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &createdUser, nil
}

// Delete deletes a User from the account.
func (u *Users) Delete(ctx context.Context, userID string) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	url, err := url.JoinPath(u.baseClient.APIUrl, "users", userID)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = u.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodDelete,
		URL:    url,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}

// ResendInvite sends a confirmation email again.
func (u *Users) ResendInvite(ctx context.Context, userID string) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	url, err := url.JoinPath(u.baseClient.APIUrl, "users", userID, "resend_invite")
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = u.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodPatch,
		URL:    url,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}

// AssignRoles adds new roles for a User with the given userID.
func (u *Users) AssignRoles(ctx context.Context, userID string, roles []Role) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	if len(roles) == 0 {
		return iamerrors.Error{Err: iamerrors.ErrUserRolesRequired, Desc: "No roles for User was provided."}
	}

	return u.manageRoles(ctx, http.MethodPut, userID, roles)
}

// UnassignRoles removes roles from a User with the given userID.
func (u *Users) UnassignRoles(ctx context.Context, userID string, roles []Role) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}
	if len(roles) == 0 {
		return iamerrors.Error{Err: iamerrors.ErrUserRolesRequired, Desc: "No roles for User was provided."}
	}

	return u.manageRoles(ctx, http.MethodDelete, userID, roles)
}

func (u *Users) manageRoles(ctx context.Context, method string, userID string, roles []Role) error {
	url, err := url.JoinPath(u.baseClient.APIUrl, "users", userID, "roles")
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	body, err := json.Marshal(manageRolesRequest{
		Roles: roles,
	})
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = u.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: method,
		URL:    url,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}
