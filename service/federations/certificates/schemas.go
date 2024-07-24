package certificates

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

type ListResponse struct {
	Certificates []Certificate `json:"certificates"`
}

type CreateResponse struct {
	Certificate
}

type GetResponse struct {
	Certificate
}

type UpdateResponse struct {
	Certificate
}

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

type UpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
