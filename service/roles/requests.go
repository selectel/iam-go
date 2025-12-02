package roles

import (
	"context"
	"net/http"
	"net/url"

	"github.com/selectel/iam-go/iamerrors"
	"github.com/selectel/iam-go/internal/client"
)

const apiVersion = "iam/v1"

// Service is used to communicate with the Roles API.
type Service struct {
	baseClient *client.BaseClient
}

// New initialises Service with the given client.
func New(baseClient *client.BaseClient) *Service {
	return &Service{
		baseClient: baseClient,
	}
}

// List returns a list of roles available for assignment.
func (s *Service) List(ctx context.Context) (*ListResponse, error) {
	path, err := url.JoinPath(apiVersion, "roles")
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

	var roles ListResponse
	err = client.UnmarshalJSON(response, &roles)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	return &roles, nil
}
