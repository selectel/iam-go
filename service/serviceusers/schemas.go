package serviceusers

import "github.com/selectel/iam-go/service/roles"

type listResponse struct {
	Users []ServiceUserListResponse `json:"users"`
}

// ServiceUserListResponse represents a Selectel Service Users in list response.
type ServiceUserListResponse struct {
	ID      string       `json:"id"`
	Enabled bool         `json:"enabled"`
	Name    string       `json:"name"`
	Roles   []roles.Role `json:"roles"`
}

// ServiceUser represents a Selectel Service User.
type ServiceUser struct {
	ID      string       `json:"id"`
	Enabled bool         `json:"enabled"`
	Name    string       `json:"name"`
	Roles   []roles.Role `json:"roles"`
	Groups  []Group      `json:"groups"`
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
	Enabled  bool
	Name     string
	Password string
	GroupIDs []string
	Roles    []roles.Role
}

type CreateResponse struct {
	ID      string       `json:"id"`
	Enabled bool         `json:"enabled"`
	Name    string       `json:"name"`
	Roles   []roles.Role `json:"roles"`
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
