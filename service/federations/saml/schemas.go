package saml

// Federation represents basic information about Federation.
type Federation struct {
	AccountID          string `json:"account_id"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Alias              string `json:"alias"`
	Issuer             string `json:"issuer"`
	SSOUrl             string `json:"sso_url"`
	SignAuthnRequests  bool   `json:"sign_authn_requests"`
	ForceAuthn         bool   `json:"force_authn"`
	SessionMaxAgeHours int    `json:"session_max_age_hours"`
	AutoUsersCreation  bool   `json:"auto_users_creation"`
	EnableGroupMapping bool   `json:"enable_group_mappings"`
}

// ListResponse represents all federations in account.
type ListResponse struct {
	Federations []Federation `json:"federations"`
}

// CreateResponse represents a configured Federation.
type CreateResponse struct {
	Federation
}

// GetResponse represents an existing Federation.
type GetResponse struct {
	Federation
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Issuer             string `json:"issuer"`
	SSOUrl             string `json:"sso_url"`
	SignAuthnRequests  bool   `json:"sign_authn_requests,omitempty"`
	ForceAuthn         bool   `json:"force_authn,omitempty"`
	SessionMaxAgeHours int    `json:"session_max_age_hours"`
	AutoUsersCreation  bool   `json:"auto_users_creation"`
	EnableGroupMapping bool   `json:"enable_group_mappings"`
}

// UpdateRequest is used to set options for Update method.
type UpdateRequest struct {
	Name               string  `json:"name,omitempty"`
	Description        *string `json:"description,omitempty"`
	Alias              string  `json:"alias,omitempty"`
	Issuer             string  `json:"issuer,omitempty"`
	SSOUrl             string  `json:"sso_url,omitempty"`
	SignAuthnRequests  *bool   `json:"sign_authn_requests,omitempty"`
	ForceAuthn         *bool   `json:"force_authn,omitempty"`
	SessionMaxAgeHours int     `json:"session_max_age_hours,omitempty"`
	AutoUsersCreation  *bool   `json:"auto_users_creation,omitempty"`
	EnableGroupMapping *bool   `json:"enable_group_mappings,omitempty"`
}

// FederationPreview represents preview information about Federation.
type FederationPreview struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Alias       string `json:"alias"`
}

// GroupMapping represents mapping between internal and external group.
type GroupMapping struct {
	InternalGroupID string `json:"internal_group_id"`
	ExternalGroupID string `json:"external_group_id"`
}

// GroupMappingsRequest is used to set options for UpdateGroupMappings method.
type GroupMappingsRequest struct {
	GroupMappings []GroupMapping `json:"group_mappings"`
}

// GroupMappingsResponse represents all mappings for the specified Federation.
type GroupMappingsResponse struct {
	GroupMappings []GroupMapping `json:"group_mappings"`
}
