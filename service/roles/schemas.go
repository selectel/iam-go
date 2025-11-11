package roles

// Name represents a role, which can be assigned to a user or a service user.
// For additional information, see
// https://docs.selectel.ru/control-panel-actions/users-and-roles/user-types-and-roles/#user-roles.
type Name string

const (
	// Account owner.
	AccountOwner Name = "account_owner"

	// User administrator.
	IAMAdmin Name = "iam.admin"

	// Account/Project administrator.
	Member Name = "member"

	// Account/Project reader.
	Reader Name = "reader"

	// Billing administrator.
	Billing Name = "billing"

	// Object storage administrator. Can be assigned only to a service user.
	ObjectStorageAdmin Name = "object_storage:admin"

	// Object storage user. Can be assigned only to a service user.
	ObjectStorageUser Name = "object_storage_user"
)

// Scope represents a scope of a role.
type Scope string

const (
	// Project scope.
	Project Scope = "project"

	// Account scope.
	Account Scope = "account"
)

type Role struct {
	ProjectID string `json:"project_id,omitempty"`
	RoleName  Name   `json:"role_name"`
	Scope     Scope  `json:"scope"`
}
