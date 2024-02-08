package testdata

const (
	TestToken = "test-token"
	TestURL   = "http://example.org/"
	TestUserAgent = "iam-go/v0.0.1"
)

const TestDoRequestRaw = `{
	"id": "test-id",
	"name": "test-name"
}`

const TestDoRequestErr = `{
	"code": "REQUEST_FORBIDDEN",
	"message": "You don't have permission to do this"
}`