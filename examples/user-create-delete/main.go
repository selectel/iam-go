package main

import (
	"context"
	"fmt"

	"github.com/selectel/iam-go"
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/users"
)

const (
	// Billing administrator.
	Billing string = "billing"

	// Account scope.
	AccountScope string = "account"
)

var (
	// KeystoneToken
	token          = "gAAAAA..."
	deleteAfterRun = false

	// Prefix to be added to User-Agent.
	postfix = "iam-go"

	// Email of the User to create.
	email = "testmail@example.com"
)

func main() {
	// Create a new IAM client.
	iamClient, err := iam.New(
		iam.WithAuthOpts(&iam.AuthOpts{KeystoneToken: token}),
		iam.WithUserAgentPostfix(postfix),
	)
	// Handle the error.
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the Users instance.
	usersAPI := iamClient.Users

	// Prepare an empty context.
	ctx := context.Background()

	// Create a new User.
	user, err := usersAPI.Create(ctx, users.CreateRequest{
		AuthType:   users.Local,
		Email:      email,
		Federation: nil,
		Roles:      []roles.Role{{Scope: AccountScope, RoleName: Billing}},
	})
	// Handle the error.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Step 1: Created User ID: %s Keystone ID: %s\n", user.ID, user.KeystoneID)

	if deleteAfterRun {
		// Delete an existing User.
		err = usersAPI.Delete(ctx, user.ID)

		// Handle the error.
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Step 2: Deleted User ID: %s", user.ID)
	}
}
