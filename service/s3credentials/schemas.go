package s3credentials

// CreateResponse represents a S3 Credentials for the given user.
// It contains "secret_key" field, which appears only once and only after creating.
type CreateResponse struct {
	Credential
	SecretKey string `json:"secret_key"`
}

// Credential represents basic information about a Credential.
type Credential struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	AccessKey string `json:"access_key"`
}

// ListResponse represents an S3 Credentials for the given user.
type ListResponse struct {
	Credentials []Credential `json:"credentials"`
}

type createRequest struct {
	Name      string `json:"name,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}
