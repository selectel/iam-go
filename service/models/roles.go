package models

// RoleName represents a role, which can be assigned to a user or service user. 
// For additional information, see https://docs.selectel.ru/control-panel-actions/users-and-roles/user-types-and-roles/#user-roles
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

	// Object storage administrator. Can be assigned only to a service user.
	ObjectStorageAdmin RoleName = "object_storage:admin"

	// Object storage user. Can be assigned only to a service user.
	ObjectStorageUser RoleName = "object_storage_user"
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
	ProjectID string   `json:"project_id,omitempty"`
	RoleName  RoleName `json:"role_name"`
	Scope     Scope    `json:"scope"`
}