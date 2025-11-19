package roles

// Role represents a scope/role pair used when managing assignments for users, groups and service users.
type Role struct {
	ProjectID string `json:"project_id,omitempty"`
	RoleName  string `json:"role_name"`
	Scope     string `json:"scope"`
}

// AvailableRole describes a role that can be assigned via IAM.
type AvailableRole struct {
	AvailableInOnboarding bool     `json:"available_in_onboarding"`
	Category              string   `json:"category"`
	Description           string   `json:"description"`
	ID                    string   `json:"id"`
	Scopes                []string `json:"scopes"`
	SubjectTypes          []string `json:"subject_types"`
}

// ListResponse is a response payload for the GET iam/v1/roles endpoint.
type ListResponse struct {
	Roles []AvailableRole `json:"roles"`
}
