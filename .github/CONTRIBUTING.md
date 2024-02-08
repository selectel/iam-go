# Contributing

Before creating a PR please create an issue that will describe a problem.

## Project structure

Every API part should be implemented in its separate package.

Any package which implements methods to work with IAM API uses the
following structure:

```
(iam_api_component)/
├── testdata
|   └── fixtures.go      # Tests fixtures
├── doc.go               # Documentation at the godoc.org
├── requests.go          # Methods to work with the API
├── requests_test.go     # Tests for all implemented requests
└── schemas.go           # Models and types
```

## Tests

Please implement tests for all methods that you're creating.

You can use: 
* [httpmock](https://github.com/jarcoal/httpmock) to mock requests and responses.
* [testify](https://github.com/stretchr/testify) to easily write assertions and validations.
