package saml

// Federation represents basic information about Federation.
type Federation struct {
	AccountID          string `json:"account_id"`
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Issuer             string `json:"issuer"`
	SSOUrl             string `json:"sso_url"`
	SignAuthnRequests  bool   `json:"sign_authn_requests"`
	ForceAuthn         bool   `json:"force_authn"`
	SessionMaxAgeHours int    `json:"session_max_age_hours"`
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
	Description        string `json:"description,omitempty"`
	Issuer             string `json:"issuer"`
	SSOUrl             string `json:"sso_url"`
	SignAuthnRequests  bool   `json:"sign_authn_requests,omitempty"`
	ForceAuthn         bool   `json:"force_authn,omitempty"`
	SessionMaxAgeHours int    `json:"session_max_age_hours"`
}

// UpdateRequest is used to set options for Update method.
type UpdateRequest struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	Issuer             string `json:"issuer,omitempty"`
	SSOUrl             string `json:"sso_url,omitempty"`
	SignAuthnRequests  *bool  `json:"sign_authn_requests,omitempty"`
	ForceAuthn         *bool  `json:"force_authn,omitempty"`
	SessionMaxAgeHours int    `json:"session_max_age_hours,omitempty"`
}
