package saml

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
	"github.com/selectel/iam-go/service/federations/saml/certificates"
)

const apiVersion = "v1"

// Service is used to communicate with the Federations API.
type Service struct {
	Certificates *certificates.Service
	baseClient   *client.BaseClient
}

// New Initialises Service with the given client.
func New(baseClient *client.BaseClient) *Service {
	return &Service{
		Certificates: certificates.New(baseClient),
		baseClient:   baseClient,
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

// Exists checks that Federation with federationID exists.
func (s *Service) Exists(ctx context.Context, federationID string) (bool, error) {
	if federationID == "" {
		return false, iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID)
	if err != nil {
		return false, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodHead,
		Path:   path,
	})
	if err != nil {
		if errors.Is(err, iamerrors.ErrFederationNotFound) {
			return false, nil
		}

		//nolint:wrapcheck // DoRequest already wraps the error.
		return false, err
	}

	return true, nil
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
	if input.SessionMaxAgeHours == 0 {
		return nil, iamerrors.Error{
			Err:  iamerrors.ErrFederationMaxAgeHoursRequired,
			Desc: "No Max Age Hours for Federation was provided.",
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

func (s *Service) getFederationResource(
	ctx context.Context, federationID string, segments []string, output interface{},
) error {
	if federationID == "" {
		return iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	pathSegments := append([]string{apiVersion, "federations", "saml", federationID}, segments...)

	path, err := url.JoinPath(pathSegments[0], pathSegments[1:]...)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	response, err := s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodGet,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	err = client.UnmarshalJSON(response, output)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	return nil
}

// Preview returns preview information of Federation using federationID or alias.
func (s *Service) Preview(ctx context.Context, federationID string) (*FederationPreview, error) {
	var preview FederationPreview
	err := s.getFederationResource(ctx, federationID, []string{"preview"}, &preview)
	if err != nil {
		return nil, err
	}

	return &preview, nil
}

// GetGroupMappings returns a list of mappings for the Federation.
func (s *Service) GetGroupMappings(ctx context.Context, federationID string) (*GroupMappingsResponse, error) {
	var mappings GroupMappingsResponse
	err := s.getFederationResource(ctx, federationID, []string{"group-mappings"}, &mappings)
	if err != nil {
		return nil, err
	}

	return &mappings, nil
}

// UpdateGroupMappings updates mappings for the Federation.
func (s *Service) UpdateGroupMappings(
	ctx context.Context, federationID string, input GroupMappingsRequest,
) (*GroupMappingsResponse, error) {
	if federationID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "saml", federationID, "group-mappings")
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

	var mappings GroupMappingsResponse
	err = client.UnmarshalJSON(response, &mappings)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	return &mappings, nil
}

func buildExternalGroupMappingPath(
	federationID, groupID, externalGroupID string,
) (string, error) {
	if federationID == "" {
		return "", iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}
	if groupID == "" {
		return "", iamerrors.Error{Err: iamerrors.ErrGroupIDRequired, Desc: "No groupID was provided."}
	}
	if externalGroupID == "" {
		return "", iamerrors.Error{Err: iamerrors.ErrInputDataRequired, Desc: "No externalGroupID was provided."}
	}

	path, err := url.JoinPath(
		apiVersion,
		"federations",
		"saml",
		federationID,
		"group-mappings",
		groupID,
		"external-groups",
		externalGroupID,
	)
	if err != nil {
		return "", iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	return path, nil
}

// AddExternalGroupMapping creates mapping between internal and external group.
func (s *Service) AddExternalGroupMapping(
	ctx context.Context, federationID, groupID, externalGroupID string,
) error {
	path, err := buildExternalGroupMappingPath(federationID, groupID, externalGroupID)
	if err != nil {
		return err
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodPut,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}

// DeleteExternalGroupMapping deletes mapping between internal and external group.
func (s *Service) DeleteExternalGroupMapping(
	ctx context.Context, federationID, groupID, externalGroupID string,
) error {
	path, err := buildExternalGroupMappingPath(federationID, groupID, externalGroupID)
	if err != nil {
		return err
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

// ExternalGroupMappingExists checks that internal and external groups are mapped.
func (s *Service) ExternalGroupMappingExists(
	ctx context.Context, federationID, groupID, externalGroupID string,
) (bool, error) {
	path, err := buildExternalGroupMappingPath(federationID, groupID, externalGroupID)
	if err != nil {
		return false, err
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   nil,
		Method: http.MethodHead,
		Path:   path,
	})
	if err != nil {
		if errors.Is(err, iamerrors.ErrFederationNotFound) ||
			errors.Is(err, iamerrors.ErrGroupNotFound) ||
			errors.Is(err, iamerrors.ErrUserOrGroupNotFound) {
			return false, nil
		}

		//nolint:wrapcheck // DoRequest already wraps the error.
		return false, err
	}

	return true, nil
}
