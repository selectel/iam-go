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
			"issuer": "https://idp.example.com",
			"client_id": "test_client_id",
			"client_secret": "test_client_secret",
			"auth_url": "https://idp.example.com/authorize",
			"token_url": "https://idp.example.com/token",
			"jwks_url": "https://idp.example.com/.well-known/jwks.json",
			"session_max_age_hours": 24
		}
	]
}`

const TestGetFederationResponse = `{
	"id": "123",
	"account_id": "123",
	"name": "test_name",
	"description": "test_description",
	"alias": "test_alias",
	"issuer": "https://idp.example.com",
	"client_id": "test_client_id",
	"client_secret": "test_client_secret",
	"auth_url": "https://idp.example.com/authorize",
	"token_url": "https://idp.example.com/token",
	"jwks_url": "https://idp.example.com/.well-known/jwks.json",
	"session_max_age_hours": 24,
	"auto_users_creation": true,
	"enable_group_mappings": true
}`

const TestCreateFederationResponse = `{
	"id": "123",
	"account_id": "123",
	"name": "test_name",
	"description": "test_description",
	"alias": "test_alias",
	"issuer": "https://idp.example.com",
	"client_id": "test_client_id",
	"client_secret": "test_client_secret",
	"auth_url": "https://idp.example.com/authorize",
	"token_url": "https://idp.example.com/token",
	"jwks_url": "https://idp.example.com/.well-known/jwks.json",
	"session_max_age_hours": 24,
	"auto_users_creation": true,
	"enable_group_mappings": true
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`

const TestFederationNotFoundErr = `{
	"code": "FEDERATION_NOT_FOUND",
	"message": "Federation not found"
}`

const TestGroupMappingsResponse = `{
	"group_mappings": [
		{
			"internal_group_id": "456",
			"external_group_id": "external-group"
		}
	]
}`

const TestUserOrGroupNotFoundErr = `{
	"code": "USER_OR_GROUP_NOT_FOUND",
	"message": "User or group not found"
}`
