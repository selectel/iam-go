package certificates

// Certificate represents a Federation Certificate.
type Certificate struct {
	ID           string `json:"id"`
	AccountID    string `json:"account_id"`
	FederationID string `json:"federation_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	NotBefore    string `json:"not_before"`
	NotAfter     string `json:"not_after"`
	Fingerprint  string `json:"fingerprint"`
	Data         string `json:"data"`
}

// ListResponse represents all certificates for the specified Federation.
type ListResponse struct {
	Certificates []Certificate `json:"certificates"`
}

// CreateResponse represents a configured Federation Certificate.
type CreateResponse struct {
	Certificate
}

// GetResponse represents an existing Federation Certificate.
type GetResponse struct {
	Certificate
}

// UpdateResponse represents an updated Federation Certificate.
type UpdateResponse struct {
	Certificate
}

// CreateRequest is used to set options for Create method.
type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

// UpdateRequest is used to set options for Update method.
type UpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
