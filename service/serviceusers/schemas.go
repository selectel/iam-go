package serviceusers

type RoleName string

const (
	// Account owner.
	AccountOwner RoleName = "account_owner"

	// User administrator.
	IAMAdmin RoleName = "iam_admin"

	// Account/Project administrator.
	Member RoleName = "member"

	// Account/Project reader.
	Reader RoleName = "reader"

	// Billing administrator.
	Billing RoleName = "billing"

	// Object storage administrator.
	ObjectStorageAdmin RoleName = "object_storage:admin"

	// Object storage user.
	ObjectStorageUser RoleName = "object_storage_user"
)

type Scope string

const (
	// Project scope.
	Project Scope = "project"

	// Account scope.
	Account Scope = "account"
)

// ServiceUser represents a Selectel Service User.
type ServiceUser struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
	Roles   []Role `json:"roles"`
}

type Role struct {
	ProjectID   string   `json:"project_id,omitempty"`
	ProjectName string   `json:"project_name,omitempty"`
	RoleName    RoleName `json:"role_name"`
	Scope       Scope    `json:"scope"`
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	Enabled  bool
	Name     string
	Password string
	Roles    []Role
}

// UpdateRequest is used to set options for Update method.
type UpdateRequest struct {
	Enabled  bool
	Name     string
	Password string
}

type createRequest struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Roles    []Role `json:"roles,omitempty"`
}

type updateRequest struct {
	Enabled  *bool  `json:"enabled,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type manageRolesRequest struct {
	Roles []Role `json:"roles"`
}

type listResponse struct {
	Users []ServiceUser `json:"users"`
}
