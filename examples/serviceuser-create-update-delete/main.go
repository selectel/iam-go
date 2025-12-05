package main

import (
	"context"
	"fmt"

	"github.com/selectel/iam-go"
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/serviceusers"
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

	// Client User-Agent prefix to be prepended to the library User-Agent.
	clientUserAgent = "iam-go"

	// Name of the Service User to create.
	name = "service-user"

	// Password of the Service User to create.
	password = "Qazwsxedc123"
)

func main() {
	// Create a new IAM client.
	iamClient, err := iam.New(
		iam.WithAuthOpts(&iam.AuthOpts{KeystoneToken: token}),
		iam.WithClientUserAgent(clientUserAgent),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the Service User instance.
	serviceUsersAPI := iamClient.ServiceUsers

	// Prepare an empty context.
	ctx := context.Background()

	// Create a new Service User.
	serviceUser, err := serviceUsersAPI.Create(ctx, serviceusers.CreateRequest{
		Enabled:  true,
		Name:     name,
		Password: password,
		Roles:    []roles.Role{{Scope: AccountScope, RoleName: Billing}},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Step 1: Created Service User ID: %s\n", serviceUser.ID)

	// Disable the just-created Service User.
	_, err = serviceUsersAPI.Update(ctx, serviceUser.ID, serviceusers.UpdateRequest{
		Enabled: false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Step 2: Disabled Service User ID %s\n", serviceUser.ID)

	// Disabled by default
	if deleteAfterRun {
		// Delete an existing Service User.
		err = serviceUsersAPI.Delete(ctx, serviceUser.ID)

		// Handle the error.
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Step 3: Deleted Service User ID %s\n", serviceUser.ID)
	}
}
