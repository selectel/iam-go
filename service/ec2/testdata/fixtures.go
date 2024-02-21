package testdata

const (
	TestToken = "test-token"
	TestURL   = "http://example.org/"
)

const TestGetCredentialsResponse = `{
	"credentials": [{
		"name": "12345",
		"project_id": "test-project",
		"access_key": "test-access-key"
	}]
}`

const TestCreateCredentialResponse = `{
	"name": "12345",
	"project_id": "test-project",
	"access_key": "test-access-key",
	"secret_key": "test-secret-key"	
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`
