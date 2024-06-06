package serviceusers

import "github.com/selectel/iam-go/service/roles"

// ListResponse represents all Service Users in account.
type ListResponse struct {
	Users []ServiceUser `json:"users"`
}

// ServiceUser represents basic information about a Selectel Service User.
type ServiceUser struct {
	ID      string       `json:"id"`
	Enabled bool         `json:"enabled"`
	Name    string       `json:"name"`
	Roles   []roles.Role `json:"roles"`
}

// GetResponse represents a Selectel Service User.
type GetResponse struct {
	ServiceUser
	Groups []Group `json:"groups"`
}

// CreateResponse represents a Selectel Service User.
type CreateResponse struct {
	ServiceUser
}

// UpdateResponse represents a Selectel Service User.
type UpdateResponse struct {
	ServiceUser
	Groups []Group `json:"groups"`
}

// Group represents information about the Group the user is a member of.
type Group struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Roles       []roles.Role `json:"roles"`
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	Enabled  bool
	Name     string
	Password string
	GroupIDs []string
	Roles    []roles.Role
}

// UpdateRequest is used to set options for Update method.
type UpdateRequest struct {
	Enabled  bool
	Name     string
	Password string
}

type createRequest struct {
	Enabled  bool         `json:"enabled,omitempty"`
	Name     string       `json:"name,omitempty"`
	Password string       `json:"password,omitempty"`
	GroupIds []string     `json:"group_ids,omitempty"`
	Roles    []roles.Role `json:"roles,omitempty"`
}

type updateRequest struct {
	Enabled  *bool  `json:"enabled,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type manageRolesRequest struct {
	Roles []roles.Role `json:"roles"`
}
