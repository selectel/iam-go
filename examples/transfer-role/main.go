package main

import (
	"context"
	"fmt"

	"github.com/selectel/iam-go"
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/users"
)

func main() {
	// KeystoneToken
	token := "gAAAAA..."

	// Prefix to be added to User-Agent.
	prefix := "iam-go"

	// ID of the User to assign role to.
	userID := "654321_65432"

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

	// Get the Users instance.
	usersAPI := iamClient.Users

	// Prepare an empty context.
	ctx := context.Background()

	// List the roles assigned to each user and find a billing.
	allUsers, err := usersAPI.List(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	var chosenUser *users.User
	for _, user := range allUsers {
		for _, role := range user.Roles {
			if role.RoleName == roles.Billing && user.ID != "account_root" {
				chosenUser = &user
				break
			}
		}
		if chosenUser != nil {
			break
		}
	}

	if chosenUser == nil {
		fmt.Println("No billing role was found")
		return
	}

	// Step 1
	fmt.Printf("Step 1: User %s with the Billing role was found\n", chosenUser.ID)

	// Unassign the role.
	err = usersAPI.UnassignRoles(
		ctx,
		chosenUser.ID,
		[]roles.Role{{Scope: roles.Account, RoleName: roles.Billing}},
	)

	// Handle the error.
	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 2
	fmt.Printf("Step 2: Unassigned the Billing role from User %s \n", chosenUser.ID)

	// Assign the role.
	err = usersAPI.AssignRoles(
		ctx,
		userID,
		[]roles.Role{{Scope: roles.Account, RoleName: roles.Billing}},
	)

	// Handle the error.
	if err != nil {
		fmt.Println(err)
	}

	// Step 3
	fmt.Printf("Step 3: Assigned the Billing role to User %s \n", userID)
}
