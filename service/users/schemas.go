package users

type AuthType string

const (
	Local     AuthType = "local"
	Federated AuthType = "federated"
)

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
)

type Scope string

const (
	// Project scope.
	Project Scope = "project"

	// Account scope.
	Account Scope = "account"
)

// User represents a Selectel Panel User.
type User struct {
	AuthType   AuthType    `json:"auth_type"`
	Federation *Federation `json:"federation,omitempty"`
	Roles      []Role      `json:"roles"`
	ID         string      `json:"id"`
	KeystoneID string      `json:"keystone_id"`
}

type Federation struct {
	ExternalID string `json:"external_id"`
	ID         string `json:"id"`
}

type Role struct {
	ProjectID string   `json:"project_id,omitempty"`
	RoleName  RoleName `json:"role_name"`
	Scope     Scope    `json:"scope"`
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	AuthType   AuthType
	Email      string
	Federation *Federation
	Roles      []Role
}

type createRequest struct {
	AuthType          AuthType    `json:"auth_type,omitempty"`
	Email             string      `json:"email,omitempty"`
	Federation        *Federation `json:"federation,omitempty"`
	Roles             []Role      `json:"roles,omitempty"`
	SubscriptionsOnly bool        `json:"subscriptions_only"` // Issue, should be hardcoded to `false`
	Subscriptions     []string    `json:"subscriptions"`      // Issue, should be hardcoded to `[]`
}

type manageRolesRequest struct {
	Roles []Role `json:"roles"`
}

type listResponse struct {
	Users []User `json:"users"`
}
