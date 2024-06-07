package groups

import (
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/users"
)

// ListResponse represents all Groups in account.
type ListResponse struct {
	Groups []Group `json:"groups"`
}

// Group represents basic information about a Group.
type Group struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Roles       []roles.Role `json:"roles"`
}

// GetResponse represents a Group of users.
type GetResponse struct {
	Group
	ServiceUsers []ServiceUser `json:"service_users"`
	Users        []User        `json:"users"`
}

// CreateResponse represents a Group of users.
type CreateResponse struct {
	Group
	ServiceUsers []ServiceUser `json:"service_users"`
	Users        []User        `json:"users"`
}

// UpdateResponse represents a Group of users.
type UpdateResponse struct {
	Group
	ServiceUsers []ServiceUser `json:"service_users"`
	Users        []User        `json:"users"`
}

// ServiceUser represents a Selectel Service User in Group.
type ServiceUser struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
}

// User represents a Selectel Panel User in Group.
type User struct {
	AuthType   users.AuthType    `json:"auth_type"`
	Federation *users.Federation `json:"federation,omitempty"`
	ID         string            `json:"id"`
	KeystoneID string            `json:"keystone_id"`
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// UpdateRequest is used as options for Update method.
type UpdateRequest struct {
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type manageRolesRequest struct {
	Roles []roles.Role `json:"roles"`
}

type manageUsersRequest struct {
	KeystoneIds []string `json:"keystone_ids"`
}
