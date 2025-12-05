# iam-go: Go SDK for IAM API

[![Go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/selectel/iam-go/)
[![Go Report Card](https://goreportcard.com/badge/github.com/selectel/iam-go)](https://goreportcard.com/report/github.com/selectel/iam-go)
[![codecov](https://codecov.io/gh/Selectel/iam-go/branch/main/graph/badge.svg)](https://codecov.io/gh/Selectel/iam-go)

Package iam-go provides Go SDK to work with the Selectel IAM API.

Jump To:
* [Documentation](#Documentation)
* [Getting started](#Getting-started)
* [Additional info](#Additional-info)

## Documentation

The Go library documentation is available at [go.dev](https://pkg.go.dev/github.com/selectel/iam-go/).

## Getting started

You can use this library to work with the following objects of the Selectel IAM API:

* [users](https://pkg.go.dev/github.com/selectel/iam-go/service/users)
* [serviceusers](https://pkg.go.dev/github.com/selectel/iam-go/service/serviceusers)
* [groups](https://pkg.go.dev/github.com/selectel/iam-go/service/groups)
* [s3credentials](https://pkg.go.dev/github.com/selectel/iam-go/service/s3credentials)
* [federations (saml)](https://pkg.go.dev/github.com/selectel/iam-go/service/federations/saml)

### Installation

You can install `iam-go` via `go get` command:

```bash
go get github.com/selectel/iam-go
```

### Authentication

To work with the Selectel IAM API you first need to:

* Create a Selectel account: [registration page](https://my.selectel.ru/registration).
* Retrieve a Keystone Token for your account via [API](https://developers.selectel.com/docs/control-panel/authorization/#obtain-keystone-token) or [go-selvpcclient](https://github.com/selectel/go-selvpcclient).

After that initialize `Client` with the retrieved token.


### Usage example

> [!NOTE] It is highly recommended to use the `WithClientUserAgent` option to set a custom User-Agent for the client.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/selectel/iam-go"
)

func main() {
    // A KeystoneToken to work with the Selectel IAM API.
    // It should be Service User Token
    token := "gAAAAABeVNzu-..."

    // A client User-Agent prefix to be prepended to the library User-Agent.
    userAgent := "iam-custom"

    // Create a new IAM client.
    iamClient, err := iam.New(
    	iam.WithAuthOpts(&iam.AuthOpts{KeystoneToken: token}),
    	iam.WithClientUserAgent(userAgent),
    )
    // Handle the error.
    if err != nil {
        log.Fatalf("Error occured: %s", err)
        return
    }

    // Get the Users instance.
    usersAPI := iamClient.Users

    // Prepare an empty context.
    ctx := context.Background()

    // Get all users.
    users := usersAPI.List(ctx)

    // Print info about each user.
    for _, user := range users {
        fmt.Println("ID:", user.ID)
        fmt.Println("KeystoneID:", user.Keystone.ID)
        fmt.Println("AuthType:", user.AuthType)
    }
}
```

## Additional info
* See [examples](./examples) for more code examples of iam-go
* Read [docs](./docs) for advanced topics and guides
