package main

import (
	"context"
	"fmt"

	iam "github.com/selectel/iam-go"
)

var (
	// KeystoneToken
	token          = "gAAAAA..."
	deleteAfterRun = false

	// Prefix to be added to User-Agent.
	prefix = "iam-go"

	// ID of the User to create S3 Credentials for.
	userID = "a1b2c3..."

	// Name of the S3 Credentials to create.
	name = "my-s3-credentials"

	// Project ID to create the S3 Credentials for.
	projectID = "a1b2c3..."
)

func main() {
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

	// Get the S3 Credentials API instance.
	s3CredAPI := iamClient.S3Credentials

	// Prepare an empty context.
	ctx := context.Background()

	// Create a new S3 Credentials for the Service User.
	credentials, err := s3CredAPI.Create(
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

	fmt.Printf("Step 1: Created credentials Secret Key: %s Access Key: %s\n", credentials.SecretKey,
		credentials.AccessKey)

	if deleteAfterRun {
		// Delete an existing S3 Credentials.
		err = s3CredAPI.Delete(ctx, userID, credentials.AccessKey)

		// Handle the error.
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Step 2: Deleted credentials")
	}
}
