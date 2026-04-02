package groupmappings

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
)

const apiVersion = "v1"

// Service is used to communicate with the OIDC Federations Group Mappings API.
type Service struct {
	baseClient *client.BaseClient
}

// New Initialises Service with the given client.
func New(baseClient *client.BaseClient) *Service {
	return &Service{
		baseClient: baseClient,
	}
}

// List returns a list of mappings for the OIDC Federation.
func (s *Service) List(ctx context.Context, federationID string) (*GroupMappingsResponse, error) {
	if federationID == "" {
		return nil, iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "oidc", federationID, "group-mappings")
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

	var mappings GroupMappingsResponse
	err = client.UnmarshalJSON(response, &mappings)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	return &mappings, nil
}

// Update updates mappings for the OIDC Federation.
func (s *Service) Update(
	ctx context.Context, federationID string, input GroupMappingsRequest,
) error {
	if federationID == "" {
		return iamerrors.Error{Err: iamerrors.ErrFederationIDRequired, Desc: "No federationID was provided."}
	}

	path, err := url.JoinPath(apiVersion, "federations", "oidc", federationID, "group-mappings")
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	body, err := json.Marshal(input)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	_, err = s.baseClient.DoRequest(ctx, client.DoRequestInput{
		Body:   bytes.NewReader(body),
		Method: http.MethodPut,
		Path:   path,
	})
	if err != nil {
		//nolint:wrapcheck // DoRequest already wraps the error.
		return err
	}

	return nil
}

// Add creates mapping between internal and external group.
func (s *Service) Add(
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

// Delete deletes mapping between internal and external group.
func (s *Service) Delete(
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

// Exists checks that internal and external groups are mapped.
func (s *Service) Exists(
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
		"oidc",
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
