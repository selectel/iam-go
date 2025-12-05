package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/selectel/iam-go/iamerrors"
)

type DoRequestInput struct {
	Body   io.Reader
	Method string
	Path   string
}

type BaseClient struct {
	// HTTPClient represents the HTTP client used to make requests.
	HTTPClient *http.Client

	// APIUrl represents a valid IAM API URL, which will be used in all requests.
	APIUrl string

	// AuthMethod contains an approach to authenticate against Selectel IAM API based on AuthOpts.
	AuthMethod AuthMethod

	// UserAgent represents a User-Agent to be added to all requests.
	UserAgent string

	// ClientUserAgent contains custom User-Agent prefix to be prepended to the library User-Agent.
	ClientUserAgent string
}

// DoRequest performs the HTTP request with the current Client.HTTPClient and User-Agent header.
//
// X-Auth-Token and other optional headers are added automatically.
func (bc *BaseClient) DoRequest(ctx context.Context, input DoRequestInput) ([]byte, error) {
	url, err := url.JoinPath(bc.APIUrl, input.Path)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	request, err := http.NewRequestWithContext(ctx, input.Method, url, input.Body)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	request.Header.Set("X-Auth-Token", bc.AuthMethod.GetKeystoneToken())
	if input.Body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	request.Header.Set("User-Agent", bc.UserAgent)

	response, err := bc.HTTPClient.Do(request)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}

	if response.StatusCode >= 400 {
		err := decodeError(response.StatusCode, body)
		return nil, err
	}

	return body, nil
}

func decodeError(statusCode int, body []byte) error {
	if statusCode == http.StatusUnauthorized {
		errDescription := string(body)
		return iamerrors.Error{Err: iamerrors.ErrAuthTokenUnathorized, Desc: errDescription}
	}

	var eg ErrorGeneric
	err := UnmarshalJSON(body, &eg)
	if err != nil {
		return iamerrors.Error{Err: iamerrors.ErrInternalAppError, Desc: err.Error()}
	}
	if e := iamerrors.GetError(eg.Code); e != nil {
		return iamerrors.Error{Err: e, Desc: eg.Message}
	}
	return iamerrors.Error{Err: iamerrors.ErrUnknown, Desc: fmt.Sprintf("%s -- %s", eg.Code, eg.Message)}
}

// ErrorGeneric represents an error returned by the IAM API.
type ErrorGeneric struct {
	// Code is a short name of error.
	Code string `json:"code"`

	// Message describes the reason of error.
	Message string `json:"message"`

	// ErrorDescription represents a human-readable description of the error.
	ErrorDescription error `json:"-"`
}

// UnmarshalJSON accepts an object in which ResposeResult.Body will be extracted.
func UnmarshalJSON(body []byte, to interface{}) error {
	err := json.Unmarshal(body, to)
	if err != nil {
		return fmt.Errorf("UnmarshalJSON() — Unmarshal error: %w", err)
	}
	return nil
}
