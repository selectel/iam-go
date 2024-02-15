package ec2

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
)

const apiVersion = "iam/v1"

// EC2 is used to communicate with the EC2(S3) API.
type EC2 struct {
	baseClient *client.BaseClient
}

// Initialises EC2 with the given client.
func New(baseClient *client.BaseClient) *EC2 {
	return &EC2{
		baseClient: baseClient,
	}
}

// List returns a list of EC2-credentials for the given user.
func (ec2 *EC2) List(ctx context.Context, userID string) ([]Credential, error) {
	if userID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "service_users", userID, "credentials")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := ec2.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var credentials listResponse
	err = client.UnmarshalJSON(response, &credentials)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return credentials.Credentials, nil
}

// Create creates a new EC2-credential for the given user.
func (ec2 *EC2) Create(ctx context.Context, userID, name, projectID string) (*CreatedCredential, error) {
	if userID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}
	if name == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrCredentialNameRequired, Desc: "No credential name was provided."}
	}
	if projectID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrProjectIDRequired, Desc: "No projectID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "service_users", userID, "credentials")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(createRequest{Name: name, ProjectID: projectID})
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := ec2.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPost,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var createdCredential CreatedCredential
	err = client.UnmarshalJSON(response, &createdCredential)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &createdCredential, nil
}

// Delete deletes an EC2-credential for the given user.
func (ec2 *EC2) Delete(ctx context.Context, userID, accessKey string) error {
	if userID == "" {
		return iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}
	if accessKey == "" {
		return iamerrors.Error{Err: iamerrors.ErrCredentialAccessKeyRequired, Desc: "No accessKey was provided."}
	}

	path, err := url.JoinPath(apiVersion, "service_users", userID, "credentials", accessKey)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	_, err = ec2.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodDelete,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}
