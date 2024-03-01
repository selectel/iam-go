package serviceusers

import "github.com/selectel/iam-go/service/roles"

// ServiceUser represents a Selectel Service User.
type ServiceUser struct {
	ID      string       `json:"id"`
	Enabled bool         `json:"enabled"`
	Name    string       `json:"name"`
	Roles   []roles.Role `json:"roles"`
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	Enabled  bool
	Name     string
	Password string
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

type listResponse struct {
	Users []ServiceUser `json:"users"`
}
