package ec2

// CreatedCredential represents a EC2 credential for the given user.
// It contains "secret_key" field, which appears only once and only after creating.
type CreatedCredential struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

// Credential represents an EC2 credential for the given user.
type Credential struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	AccessKey string `json:"access_key"`
}

type createRequest struct {
	Name      string `json:"name,omitempty"`
	ProjectID string `json:"project_id,omitempty"`
}

type listResponse struct {
	Credentials []Credential `json:"credentials"`
}
