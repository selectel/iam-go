package federations

type Federation struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Issuer             string `json:"issuer"`
	SSOUrl             string `json:"sso_url"`
	SignAuthnRequests  bool   `json:"sign_authn_requests"`
	ForceAuthn         bool   `json:"force_authn"`
	SessionMaxAgeHours int    `json:"session_max_age_hours"`
}

type FederationWithIDs struct {
	Federation
	AccountID string `json:"account_id"`
	ID        string `json:"id"`
}

type ListResponse struct {
	Federations []FederationWithIDs `json:"federations"`
}

type CreateResponse struct {
	FederationWithIDs
}

type GetResponse struct {
	FederationWithIDs
}

type CreateRequest struct {
	Federation
}

type UpdateRequest struct {
	Federation
}
