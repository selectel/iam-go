package groups

import (
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/users"
)

type listResponse struct {
	Groups []GroupListResponse `json:"groups"`
}

// GroupListResponse represents a Group in list response.
type GroupListResponse struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Roles       []roles.Role `json:"roles"`
}

// Group represents a Group of users.
type Group struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Roles        []roles.Role  `json:"roles"`
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

// ModifyRequest is used as options for Update method.
type ModifyRequest struct {
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type manageRolesRequest struct {
	Roles []roles.Role `json:"roles"`
}

type manageUsersRequest struct {
	KeystoneIds []string `json:"keystone_ids"`
}
