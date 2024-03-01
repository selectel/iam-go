package main

import (
	"context"
	"fmt"

	"github.com/selectel/iam-go"
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/serviceusers"
)

func main() {
	// KeystoneToken
	token := "gAAAAA..."

	// Prefix to be added to User-Agent.
	prefix := "iam-go"

	// Name of the Service User to create.
	name := "service-user"

	// Password of the Service User to create.
	password := "Qazwsxedc123"

	// Create a new IAM client.
	iamClient, err := iam.New(
		iam.WithAuthOpts(&iam.AuthOpts{KeystoneToken: token}),
		iam.WithUserAgentPrefix(prefix),
	)
	// Handle the error.
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
		Roles:    []roles.Role{{Scope: roles.Account, RoleName: roles.Billing}},
	})
	// Handle the error.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Step 1: Created Service User ID: %s\n", serviceUser.ID)

	// Disable the just-created Service User.
	_, err = serviceUsersAPI.Update(ctx, serviceUser.ID, serviceusers.UpdateRequest{
		Enabled: false,
	})
	// Handle the error.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Step 2: Disabled Service User ID %s\n", serviceUser.ID)

	// // Delete an existing Service User.
	// err = serviceUsersAPI.Delete(ctx, serviceUser.ID)

	// // Handle the error.
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Step 3: Deleted Service User ID %s\n", serviceUser.ID)
}
