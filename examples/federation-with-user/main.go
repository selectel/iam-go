package main

import (
	"context"
	"fmt"
	"os"

	"github.com/selectel/iam-go"
	"github.com/selectel/iam-go/service/federations/saml"
	"github.com/selectel/iam-go/service/federations/saml/certificates"
	"github.com/selectel/iam-go/service/roles"
	"github.com/selectel/iam-go/service/users"
)

const (
	// Account/Project reader.
	Reader string = "reader"

	// Account scope.
	AccountScope string = "account"
)

var (
	// KeystoneToken
	token          = "gAAAAA..."
	deleteAfterRun = false

	// Client User-Agent to be added.
	clientUserAgent = "iam-go"

	federationName               = "federation_name"
	federationDescription        = "federation_description"
	updatedFederationName        = "new_federation_name"
	updatedFederationDescription = "new_federation_description"

	certificateName        = "certificate name"
	certificateDescription = "certificate description"
	certificateFileName    = "cert.crt"

	userEmail      = "testmail@example.com"
	userExternalID = "some_id"
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

	federationsAPI := iamClient.SAMLFederations
	federationsCertificatesAPI := federationsAPI.Certificates
	usersAPI := iamClient.Users

	ctx := context.Background()

	federation, err := federationsAPI.Create(ctx, saml.CreateRequest{
		Name:               federationName,
		Description:        federationDescription,
		Issuer:             "http://localhost:8080/realms/master",
		SSOUrl:             "http://localhost:8080/realms/master/protocol/saml",
		SessionMaxAgeHours: 24,
		SignAuthnRequests:  true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Step 1: Created Federation Name: %s ID: %s\n", federation.Name, federation.ID)

	cert, err := os.ReadFile(certificateFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	certificate, err := federationsCertificatesAPI.Create(ctx, federation.ID, certificates.CreateRequest{
		Name:        certificateName,
		Description: certificateDescription,
		Data:        string(cert),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Step 2: Created Certificate for Federation ID: %s Federation ID: %s\n", certificate.ID, federation.ID)

	user, err := usersAPI.Create(ctx, users.CreateRequest{
		AuthType: users.Federated,
		Email:    userEmail,
		Federation: &users.Federation{
			ExternalID: userExternalID,
			ID:         federation.ID,
		},
		Roles: []roles.Role{{Scope: AccountScope, RoleName: Reader}},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Step 3: Created federated User ID: %s Keystone ID: %s\n", user.ID, user.KeystoneID)

	err = federationsAPI.Update(ctx, federation.ID, saml.UpdateRequest{
		Name:        updatedFederationName,
		Description: &updatedFederationDescription,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Step 4: Updated Federation Name and Description")

	if deleteAfterRun {
		// Removing User and Federation Certificate is unnecessary because removal of Federation
		// also deletes its Certificate and all attached Users
		fmt.Printf("Step 5: Deleting Federation with ID: %s\n", federation.ID)
		if err = federationsAPI.Delete(ctx, federation.ID); err != nil {
			fmt.Println(err)
		}
	}
}
