package iamerrors

import (
	"errors"
	"fmt"
)

var (
	ErrClientNoAuthOpts     = errors.New("CLIENT_NO_AUTH_METHOD")
	ErrAuthTokenUnathorized = errors.New("AUTH_TOKEN_UNAUTHORIZED")

	ErrUserNotFound           = errors.New("USER_NOT_FOUND")
	ErrDomainNotFound         = errors.New("DOMAIN_NOT_FOUND")
	ErrProjectNotFound        = errors.New("PROJECT_NOT_FOUND")
	ErrUserAlreadyExists      = errors.New("USER_ALREADY_EXISTS")
	ErrRequestValidationError = errors.New("REQUEST_VALIDATION_FAILED")
	ErrForbidden              = errors.New("REQUEST_FORBIDDEN")
	ErrUnauthorized           = errors.New("USER_UNAUTHORIZED")
	ErrInternalServerError    = errors.New("INTERNAL_SERVER_ERROR")
	ErrCredentialNotFound     = errors.New("CRED_NOT_FOUND")

	ErrUserIDRequired    = errors.New("USER_ID_REQUIRED")
	ErrProjectIDRequired = errors.New("PROJECT_ID_REQUIRED")

	ErrCredentialNameRequired      = errors.New("CREDENTIAL_NAME_REQUIRED")
	ErrCredentialAccessKeyRequired = errors.New("CREDENTIAL_ACCESS_KEY_REQUIRED")

	ErrServiceUserNameRequired     = errors.New("SERVICE_USER_NAME_REQUIRED")
	ErrServiceUserPasswordRequired = errors.New("SERVICE_USER_PASSWORD_REQUIRED")
	ErrServiceUserRolesRequired    = errors.New("SERVICE_USER_ROLES_REQUIRED")

	ErrUserRolesRequired = errors.New("USER_ROLES_REQUIRED")
	ErrUserEmailRequired = errors.New("USER_EMAIL_REQUIRED")

	ErrInputDataRequired = errors.New("INPUT_DATA_REQUIRED")

	ErrInternalAppError = errors.New("INTERNAL_APP_ERROR")

	ErrUnknown = errors.New("UNKNOWN_ERROR")

	//nolint:gochecknoglobals // stringToError is not global.
	stringToError = map[string]error{
		ErrUserNotFound.Error():                ErrUserNotFound,
		ErrClientNoAuthOpts.Error():            ErrClientNoAuthOpts,
		ErrAuthTokenUnathorized.Error():        ErrAuthTokenUnathorized,
		ErrDomainNotFound.Error():              ErrDomainNotFound,
		ErrCredentialNotFound.Error():          ErrCredentialNotFound,
		ErrProjectNotFound.Error():             ErrProjectNotFound,
		ErrUserAlreadyExists.Error():           ErrUserAlreadyExists,
		ErrRequestValidationError.Error():      ErrRequestValidationError,
		ErrForbidden.Error():                   ErrForbidden,
		ErrUnauthorized.Error():                ErrUnauthorized,
		ErrInternalServerError.Error():         ErrInternalServerError,
		ErrCredentialNameRequired.Error():      ErrCredentialNameRequired,
		ErrCredentialAccessKeyRequired.Error(): ErrCredentialAccessKeyRequired,
		ErrUserIDRequired.Error():              ErrUserIDRequired,
		ErrProjectIDRequired.Error():           ErrProjectIDRequired,
		ErrServiceUserNameRequired.Error():     ErrServiceUserNameRequired,
		ErrServiceUserPasswordRequired.Error(): ErrServiceUserPasswordRequired,
		ErrServiceUserRolesRequired.Error():    ErrServiceUserRolesRequired,
		ErrUserRolesRequired.Error():           ErrUserRolesRequired,
		ErrUserEmailRequired.Error():           ErrUserEmailRequired,
		ErrInputDataRequired.Error():           ErrInputDataRequired,
		ErrInternalAppError.Error():            ErrInternalAppError,
		ErrUnknown.Error():                     ErrUnknown,
	}
)

func GetError(errorString string) error {
	err, ok := stringToError[errorString]
	if !ok {
		return nil
	}
	return err
}

// Error represents an error returned by the IAM API. It contains a human-readable description of the error.
type Error struct {
	Err  error
	Desc string
}

func (e Error) Error() string {
	return fmt.Sprintf("iam-go: error — %s: %s", e.Err.Error(), e.Desc)
}

func (e Error) Is(err error) bool {
	return errors.Is(e.Err, err)
}