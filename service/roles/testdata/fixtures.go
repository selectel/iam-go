package testdata

const (
	TestToken = "test-token"
	TestURL   = "http://example.org/"
)

const TestListRolesResponse = `{
	"roles": [
		{
			"available_in_onboarding": true,
			"category": "general",
			"description": "Test role",
			"id": "role-id",
			"scopes": ["account"],
			"subject_types": ["user"],
			"deprecated": false
		}
	]
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`
