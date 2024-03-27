package s3credentials

// CreatedCredentials represents a S3 Credentials for the given user.
// It contains "secret_key" field, which appears only once and only after creating.
type CreatedCredentials struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

// Credentials represents an S3 Credentials for the given user.
type Credentials struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	AccessKey string `json:"access_key"`
}

type createRequest struct {
	Name      string `json:"name,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}

type listResponse struct {
	Credentials []Credentials `json:"credentials"`
}
