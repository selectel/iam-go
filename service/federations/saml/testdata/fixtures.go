package testdata

const (
	TestToken = "test-token"
	TestURL   = "http://example.org/"
)

const TestListFederationsResponse = `{
	"federations": [
		{
			"id": "123",
			"account_id": "123",
			"name": "test_name",
			"description": "test_description",
			"issuer": "test_issuer",
			"sso_url": "test_sso_url",
			"sign_authn_requests": true,
			"force_authn": true,
			"session_max_age_hours": 1
		}
	]
}`

const TestGetFederationResponse = `{
	"id": "123",
	"account_id": "123",
	"name": "test_name",
	"description": "test_description",
	"issuer": "test_issuer",
	"sso_url": "test_sso_url",
	"sign_authn_requests": true,
	"force_authn": true,
	"session_max_age_hours": 1
}`

const TestCreateFederationResponse = `{
	"id": "123",
	"account_id": "123",
	"name": "test_name",
	"description": "test_description",
	"issuer": "test_issuer",
	"sso_url": "test_sso_url",
	"sign_authn_requests": true,
	"force_authn": true,
	"session_max_age_hours": 1
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`
