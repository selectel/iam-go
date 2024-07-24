package certificates

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

// List returns a list of Certificates for the Federation.
func (s *Service) List(ctx context.Context, federationID string) (*ListResponse, error) {
	if federationID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID, "certificates")
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

	var certificates ListResponse
	err = client.UnmarshalJSON(response, &certificates)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	return &certificates, nil
}

// Get returns an info of Certificate with certificateID.
func (s *Service) Get(ctx context.Context, federationID, certificateID string) (*GetResponse, error) {
	if federationID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}
	if certificateID == "" {
		return nil, iamerrors.Error{
			Err:  iamerrors.ErrFederationCertificateIDRequired,
			Desc: "No certificateID was provided.",
		}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID, "certificates", certificateID)
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

	var certificate GetResponse
	err = client.UnmarshalJSON(response, &certificate)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	return &certificate, nil
}

// Create creates a new Certificate for the Federation.
func (s *Service) Create(ctx context.Context, federationID string, input CreateRequest) (*CreateResponse, error) {
	if federationID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID, "certificates")
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

	var createdCertificate CreateResponse
	err = client.UnmarshalJSON(response, &createdCertificate)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	return &createdCertificate, nil
}

// Update updates the Certificate with certificateID.
func (s *Service) Update(
	ctx context.Context, federationID, certificateID string, input UpdateRequest,
) (*UpdateResponse, error) {
	if federationID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}
	if certificateID == "" {
		return nil, iamerrors.Error{
			Err:  iamerrors.ErrFederationCertificateIDRequired,
			Desc: "No certificateID was provided.",
		}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID, "certificates", certificateID)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(input)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPatch,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return nil, err
	}

	var updatedCertificate UpdateResponse
	err = client.UnmarshalJSON(response, &updatedCertificate)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	return &updatedCertificate, nil
}

// Delete deletes the Certificate with certificateID.
func (s *Service) Delete(ctx context.Context, federationID, certificateID string) error {
	if federationID == "" {
		return iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}
	if certificateID == "" {
		return iamerrors.Error{
			Err:  iamerrors.ErrFederationCertificateIDRequired,
			Desc: "No certificateID was provided.",
		}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID, "certificates", certificateID)
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
