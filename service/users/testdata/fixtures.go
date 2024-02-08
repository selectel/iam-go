package testdata

const (
	TestToken = "test-token"
	TestURL   = "http://example.org"
)

const TestGetUsersResponse = `{
	"users": [
		{
			"auth_type": "local",
			"id": "123",
			"keystone_id": "123",
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
	"auth_type": "local",
	"id": "123",
	"keystone_id": "123",
	"roles": [
        {
            "scope": "account",
            "role_name": "member"
        }
    ]
}`

const TestCreateUserResponse = `{
	"auth_type": "federated",
	"keystone_id": "123",
	"id": "123",
	"roles": [
		{
			"scope": "account",
			"role_name": "member"
		}
	],
	"federation": {
		"external_id": "123",
		"id": "123"
	}
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`
