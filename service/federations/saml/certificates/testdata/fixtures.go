package testdata

const (
	TestToken = "test-token"
	TestURL   = "http://example.org/"
)

const TestListCertificatesResponse = `{
	"certificates": [
		{
			"id": "123",
			"account_id": "123",
		    "federation_id": "123",
		    "name": "test_name",
		    "description": "test_description",
		    "not_before": "2021-01-01T00:00:00Z",
		    "not_after": "2022-01-01T00:00:00Z",
		    "fingerprint": "test_fingerprint",
		    "data": "test_data"
		}
	]
}`

const TestGetCertificateResponse = `{
	"id": "123",
	"account_id": "123",
	"federation_id": "123",
	"name": "test_name",
	"description": "test_description",
	"not_before": "2021-01-01T00:00:00Z",
	"not_after": "2022-01-01T00:00:00Z",
	"fingerprint": "test_fingerprint",
	"data": "test_data"
}`

const TestCreateCertificateResponse = `{
	"id": "123",
	"account_id": "123",
	"federation_id": "123",
	"name": "test_name",
	"description": "test_description",
	"not_before": "2021-01-01T00:00:00Z",
	"not_after": "2022-01-01T00:00:00Z",
	"fingerprint": "test_fingerprint",
	"data": "test_data"
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`
