package main

import (
	"context"
	"fmt"

	iam "github.com/selectel/iam-go"
)

func main() {
	// KeystoneToken
	token := "gAAAAA..."

	// Prefix to be added to User-Agent.
	prefix := "iam-go"

	// ID of the User to create EC2 credential for.
	userID := "a1b2c3..."

	// Name of the EC2 credential to create.
	name := "my-ec2-credential"

	// Project ID to create the EC2 credential for.
	projectID := "a1b2c3..."

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

	// Get the EC2 Credentials APIinstance.
	ec2CredAPI := iamClient.EC2

	// Prepare an empty context.
	ctx := context.Background()

	// Create a new EC2 credential for the Service User.
	credential, err := ec2CredAPI.Create(
		ctx,
		userID,
		name,
		projectID,
	)
	// Handle the error.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Step 1: Created credential Secret Key: %s Access Key: %s\n", credential.SecretKey, credential.AccessKey)

	// // Delete an existing EC2 credential.
	// err = ec2CredAPI.Delete(ctx, &ec2.DeleteInput{
	// 	UserID:    userID,
	// 	AccessKey: credential.AccessKey,
	// })

	// // Handle the error.
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Step 2: Deleted credential")
}
