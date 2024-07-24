package federations

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
)

const apiVersion = "v1"

// Service is used to communicate with the Federations API.
type Service struct {
	baseClient *client.BaseClient
}

// New Initialises Service with the given client.
func New(baseClient *client.BaseClient) *Service {
	return &Service{
		baseClient: baseClient,
	}
}

// List returns a list of Federations for the account.
func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	path, err := url.JoinPath(apiVersion, "federations", "saml")
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

	var federations ListResponse
	err = client.UnmarshalJSON(response, &federations)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &federations, nil
}

// Get returns an info of Federation with federationID.
func (s *Service) Get(ctx context.Context, federationID string) (*GetResponse, error) {
	if federationID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID)
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

	var federation GetResponse
	err = client.UnmarshalJSON(response, &federation)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &federation, nil
}

// Create creates a new Federation.
func (s *Service) Create(ctx context.Context, input CreateRequest) (*CreateResponse, error) {
	if input.Name == "" {
		return nil, iamerrors.Error{
			Err:  iamerrors.ErrFederationNameRequired,
			Desc: "No Name for Federation was provided.",
		}
	}
	if input.Issuer == "" {
		return nil, iamerrors.Error{
			Err:  iamerrors.ErrFederationIssuerRequired,
			Desc: "No Issuer for Federation was provided.",
		}
	}
	if input.SSOUrl == "" {
		return nil, iamerrors.Error{
			Err:  iamerrors.ErrFederationSSOURLRequired,
			Desc: "No SSO URL for Federation was provided.",
		}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml")
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(input)
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

	var federation CreateResponse
	err = client.UnmarshalJSON(response, &federation)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &federation, nil
}

// Check checks if Federation with federationID exists.
func (s *Service) Check(ctx context.Context, federationID string) error {
	if federationID == "" {
		return iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodHead,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}

// Update updates existing Federation.
func (s *Service) Update(ctx context.Context, federationID string, input UpdateRequest) error {
	if federationID == "" {
		return iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(input)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPatch,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}

// Delete deletes a Federation from the account.
func (s *Service) Delete(ctx context.Context, federationID string) error {
	if federationID == "" {
		return iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID)
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
