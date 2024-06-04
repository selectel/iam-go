package users

import (
	"github.com/selectel/iam-go/service/roles"
)

// AuthType represents a type of authentication of a User.
type AuthType string

const (
	// Local authentication. Set by default.
	Local AuthType = "local"

	// Federated authentication. If set, the `federation` field is also has to be specified.
	Federated AuthType = "federated"
)

type listResponse struct {
	Users []UserListResponse `json:"users"`
}

// UserListResponse represents a Selectel User in list response.
type UserListResponse struct {
	AuthType   AuthType     `json:"auth_type"`
	Federation *Federation  `json:"federation,omitempty"`
	Roles      []roles.Role `json:"roles"`
	ID         string       `json:"id"`
	KeystoneID string       `json:"keystone_id"`
}

// User represents a Selectel Panel User.
type User struct {
	AuthType   AuthType     `json:"auth_type"`
	Federation *Federation  `json:"federation,omitempty"`
	Roles      []roles.Role `json:"roles"`
	ID         string       `json:"id"`
	KeystoneID string       `json:"keystone_id"`
	Groups     []Group      `json:"groups"`
}

type Federation struct {
	// ExternalID is user id that will be sent by the identity provider.
	ExternalID string `json:"external_id"`
	// ID represents identifier of federation in Selectel
	ID string `json:"id"`
}

// Group represents a Group for users.
type Group struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Roles       []roles.Role `json:"roles"`
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	AuthType   AuthType
	Email      string
	Federation *Federation
	Roles      []roles.Role
	GroupIDs   []string
}

type createRequest struct {
	AuthType          AuthType     `json:"auth_type,omitempty"`
	Email             string       `json:"email,omitempty"`
	Federation        *Federation  `json:"federation,omitempty"`
	Roles             []roles.Role `json:"roles,omitempty"`
	GroupIds          []string     `json:"group_ids,omitempty"`
	SubscriptionsOnly bool         `json:"subscriptions_only"` // Issue, should be hardcoded to `false`
	Subscriptions     []string     `json:"subscriptions"`      // Issue, should be hardcoded to `[]`
}

type CreateResponse struct {
	AuthType   AuthType     `json:"auth_type"`
	Federation *Federation  `json:"federation,omitempty"`
	Roles      []roles.Role `json:"roles"`
	ID         string       `json:"id"`
	KeystoneID string       `json:"keystone_id"`
}

type manageRolesRequest struct {
	Roles []roles.Role `json:"roles"`
}
