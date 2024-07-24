package federations

// Federation represents basic information about Federation.
type Federation struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Issuer             string `json:"issuer"`
	SSOUrl             string `json:"sso_url"`
	SignAuthnRequests  bool   `json:"sign_authn_requests"`
	ForceAuthn         bool   `json:"force_authn"`
	SessionMaxAgeHours int    `json:"session_max_age_hours,omitempty"`
}

// FederationWithIDs extends basic Federation with additional info.
type FederationWithIDs struct {
	Federation
	AccountID string `json:"account_id"`
	ID        string `json:"id"`
}

// ListResponse represents all federations in account.
type ListResponse struct {
	Federations []FederationWithIDs `json:"federations"`
}

// CreateResponse represents a configured Federation.
type CreateResponse struct {
	FederationWithIDs
}

// GetResponse represents an existing Federation.
type GetResponse struct {
	FederationWithIDs
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	Federation
}

// UpdateRequest is used to set options for Update method.
type UpdateRequest struct {
	Federation
}
