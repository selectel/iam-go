package testdata

const (
	TestToken = "test-token"
	TestURL   = "http://example.org"
)

const TestGetUsersResponse = `{
    "users": [
        {
            "name": "test",
            "enabled": true,
            "id": "123",
            "roles": [
                {
                    "scope": "account",
                    "role_name": "member"
                }
            ]
        }
	]
}`

const TestGetUserResponse = `{
	"name": "test",
	"enabled": true,
	"id": "123",
	"roles": [
		{
			"scope": "account",
			"role_name": "member"
		}
	]
}`

const TestCreateUserResponse = `{
	"name": "test",
	"enabled": true,
	"id": "123",
	"roles": [
		{
			"scope": "account",
			"role_name": "member"
		}
	]
}`

const TestUpdateUserResponse = `{
	"name": "test1",
	"enabled": true,
	"id": "123"
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`

// nolint gosec complains
const TestCreateUserInsecurePasswordErr = `{
	"code": "REQUEST_VALIDATION_FAILED",
	"message": "insecure_password"
}`
