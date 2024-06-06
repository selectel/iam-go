package testdata

const (
	TestToken = "test-token"
	TestURL   = "http://example.org/"
)

const TestListGroupsResponse = `{
	"groups": [
		{
			"id": "123",
			"name": "test_name",
			"description": "test_description",
			"roles": [
				{
					"scope": "account",
					"role_name": "member"
				}
			]
		}
	]
}`

const TestGetGroupResponse = `{
	"id": "123",
	"name": "test_name",
	"description": "test_description",
	"roles": [
				{
					"scope": "account",
					"role_name": "member"
				}
			],
	"users": [
		{
			"auth_type": "federated",
			"federation": {
				"id": "674870b5-6ad1-478e-8384-2527507fd85d",
				"external_id": "asdfasdf"
			},
			"id": "999000_33333",
			"keystone_id": "2b573185f23a40c88cbb59ed74dba683",
			"roles": null
		}
	],
	"service_users": [
		{
			"name": "1234",
			"enabled": false,
			"id": "c1f50a57fc95438aafe1e2fe87a781c2",
			"roles": null
		}
	]
}`

const TestCreateGroupResponse = `{
	"id": "123",
	"name": "test_name",
	"description": "test_description",
	"roles": [],
	"users": [],
	"service_users": []
}`

const TestUpdateGroupResponse = `{
	"id": "123",
	"name": "test_name",
	"description": "test_description",
	"roles":  [
				{
					"scope": "account",
					"role_name": "member"
				}
			],
	"users": [
		{
			"auth_type": "federated",
			"federation": {
				"id": "674870b5-6ad1-478e-8384-2527507fd85d",
				"external_id": "asdfasdf"
			},
			"id": "999000_33333",
			"keystone_id": "2b573185f23a40c88cbb59ed74dba683",
			"roles": null
		}
	],
	"service_users": [
		{
			"name": "1234",
			"enabled": false,
			"id": "c1f50a57fc95438aafe1e2fe87a781c2",
			"roles": null
		}
	]
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`
