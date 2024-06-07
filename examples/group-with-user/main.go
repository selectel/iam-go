package main

import (
	"context"
	"fmt"
	"log"

	"github.com/selectel/iam-go"
	"github.com/selectel/iam-go/service/groups"
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/users"
)

var (
	// KeystoneToken
	token          = "gAAAAA..."
	deleteAfterRun = false

	// Prefix to be added to User-Agent.
	prefix = "iam-go"

	groupName          = "test_group_name"
	description        = "group_description"
	updatedGroupName   = "new_test_group_name"
	updatedDescription = "new_group_description"
	email              = "testmail@example.com"
)

func main() {
	// Create a new IAM client.
	iamClient, err := iam.New(
		iam.WithAuthOpts(&iam.AuthOpts{KeystoneToken: token}),
		iam.WithUserAgentPrefix(prefix),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	usersAPI := iamClient.Users
	groupsAPI := iamClient.Groups
	ctx := context.Background()

	group, err := groupsAPI.Create(ctx, groups.CreateRequest{Name: groupName, Description: description})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Step 1: Created Group Name: %s ID: %s\n", group.Name, group.ID)

	user, err := usersAPI.Create(ctx, users.CreateRequest{
		AuthType:   users.Local,
		Email:      email,
		Federation: nil,
		Roles:      []roles.Role{{Scope: roles.Account, RoleName: roles.Reader}},
		GroupIDs:   []string{group.ID},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Step 2: Created User ID: %s Keystone ID: %s\n", user.ID, user.KeystoneID)

	err = groupsAPI.AssignRoles(ctx, group.ID, []roles.Role{{Scope: roles.Account, RoleName: roles.Member}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Step 3: Assigned Role %s with scope %s to Group ID: %s\n", roles.Member, roles.Account, group.ID)

	updatedGroup, err := groupsAPI.Update(ctx, group.ID, groups.UpdateRequest{Name: updatedGroupName,
		Description: &updatedDescription})
	if err != nil {
		fmt.Println(err)
		return
	}
	group.Group = updatedGroup.Group
	fmt.Printf("Step 4: Group Name and Description updated to: %s and %s\n", group.Name, group.Description)

	if deleteAfterRun {
		fmt.Printf("Step 5: Deleting Group with ID: %s\n", group.ID)
		if err = groupsAPI.Delete(ctx, group.ID); err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Step 6: Deleting User with ID: %s\n", user.ID)
		if err = usersAPI.Delete(ctx, user.ID); err != nil {
			fmt.Println(err)
		}
	}
}
