package roles

type Role struct {
	ProjectID string `json:"project_id,omitempty"`
	RoleName  string   `json:"role_name"`
	Scope     string  `json:"scope"`
}
