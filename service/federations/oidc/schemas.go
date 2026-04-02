package oidc

// Federation represents basic information about an OIDC Federation.
type Federation struct {
	ID                 string `json:"id"`
	AccountID          string `json:"account_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Alias              string `json:"alias"`
	Issuer             string `json:"issuer"`
	ClientID           string `json:"client_id"`
	ClientSecret       string `json:"client_secret"`
	AuthURL            string `json:"auth_url"`
	TokenURL           string `json:"token_url"`
	JWKSURL            string `json:"jwks_url"` //nolint:tagliatelle
	SessionMaxAgeHours int    `json:"session_max_age_hours"`
	AutoUsersCreation  bool   `json:"auto_users_creation"`
	EnableGroupMapping bool   `json:"enable_group_mappings"` //nolint:tagliatelle
}

// ListResponse represents all OIDC federations in account.
type ListResponse struct {
	Federations []Federation `json:"federations"`
}

// CreateResponse represents a configured OIDC Federation.
type CreateResponse struct {
	Federation
}

// GetResponse represents an existing OIDC Federation.
type GetResponse struct {
	Federation
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	Alias              string `json:"alias,omitempty"`
	Issuer             string `json:"issuer"`
	ClientID           string `json:"client_id"`
	ClientSecret       string `json:"client_secret"`
	AuthURL            string `json:"auth_url"`
	TokenURL           string `json:"token_url"`
	JWKSURL            string `json:"jwks_url"` //nolint:tagliatelle
	SessionMaxAgeHours int    `json:"session_max_age_hours"`
	AutoUsersCreation  bool   `json:"auto_users_creation"`
	EnableGroupMapping bool   `json:"enable_group_mappings"` //nolint:tagliatelle
}

// UpdateRequest is used to set options for Update method.
type UpdateRequest struct {
	Name               string  `json:"name,omitempty"`
	Description        *string `json:"description,omitempty"`
	Alias              string  `json:"alias,omitempty"`
	Issuer             string  `json:"issuer"`
	ClientID           string  `json:"client_id,omitempty"`
	ClientSecret       string  `json:"client_secret,omitempty"`
	AuthURL            string  `json:"auth_url,omitempty"`
	TokenURL           string  `json:"token_url,omitempty"`
	JWKSURL            string  `json:"jwks_url,omitempty"` //nolint:tagliatelle
	SessionMaxAgeHours int     `json:"session_max_age_hours,omitempty"`
	AutoUsersCreation  *bool   `json:"auto_users_creation,omitempty"`
	EnableGroupMapping *bool   `json:"enable_group_mappings,omitempty"` //nolint:tagliatelle
}
