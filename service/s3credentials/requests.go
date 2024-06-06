package s3credentials

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

// Service is used to communicate with the S3 Credentials API.
type Service struct {
	baseClient *client.BaseClient
}

// New initialises Service instance with the given client.
func New(baseClient *client.BaseClient) *Service {
	return &Service{
		baseClient: baseClient,
	}
}

// List returns a list of S3 Credentials for the given user.
func (s *Service) List(ctx context.Context, userID string) (*ListResponse, error) {
	if userID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "service_users", userID, "credentials")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var credentials ListResponse
	err = client.UnmarshalJSON(response, &credentials)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &credentials, nil
}

// Create creates a new S3 Credentials for the given user.
func (s *Service) Create(ctx context.Context, userID, name, projectID string) (*CreateResponse, error) {
	if userID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrUserIDRequired, Desc: "No userID was provided."}
	}
	if name == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrCredentialNameRequired, Desc: "No credentials name was provided."}
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

	response, err := s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPost,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var createdCredential CreateResponse
	err = client.UnmarshalJSON(response, &createdCredential)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &createdCredential, nil
}

// Delete deletes an S3 Credentials for the given user.
func (s *Service) Delete(ctx context.Context, userID, accessKey string) error {
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
	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
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
