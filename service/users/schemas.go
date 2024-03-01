package users

import "github.com/selectel/iam-go/service/models"

// AuthType represents a type of authentication of a User.
type AuthType string

const (
	// Local authentication. Set by default.
	Local AuthType = "local"

	// Federated authentication. If set, the `federation` field is also has to be specified.
	Federated AuthType = "federated"
)

// User represents a Selectel Panel User.
type User struct {
	AuthType   AuthType       `json:"auth_type"`
	Federation *Federation    `json:"federation,omitempty"`
	Roles      []models.Role `json:"roles"`
	ID         string         `json:"id"`
	KeystoneID string         `json:"keystone_id"`
}

type Federation struct {
	ExternalID string `json:"external_id"`
	ID         string `json:"id"`
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	AuthType   AuthType
	Email      string
	Federation *Federation
	Roles      []models.Role
}

type createRequest struct {
	AuthType          AuthType       `json:"auth_type,omitempty"`
	Email             string         `json:"email,omitempty"`
	Federation        *Federation    `json:"federation,omitempty"`
	Roles             []models.Role `json:"roles,omitempty"`
	SubscriptionsOnly bool           `json:"subscriptions_only"` // Issue, should be hardcoded to `false`
	Subscriptions     []string       `json:"subscriptions"`      // Issue, should be hardcoded to `[]`
}

type manageRolesRequest struct {
	Roles []models.Role `json:"roles"`
}

type listResponse struct {
	Users []User `json:"users"`
}
