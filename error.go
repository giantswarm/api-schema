package apischema

import (
	"strings"
)

// Error type returned when checking responses
// that contain an error.
type ResponseError struct {
	// High level status of the response, api_schema.go
	statusCode int
	// Project specific error response, see project that uses it
	errorCode int
	// Human readable (but not beautified) error text
	statusText string
}

// NewResponseError creates a new error based on the given response.
func NewResponseError(resp *Response) error {
	return &ResponseError{
		statusCode: resp.StatusCode,
		errorCode:  resp.ErrorCode,
		statusText: resp.StatusText,
	}
}

// Gets the status text contained in the given error
func (err *ResponseError) Error() string {
	return err.statusText
}

// Gets the project specific error code contained in the given error
func (err *ResponseError) ErrorCode() int {
	return err.errorCode
}

// Gets the high level status code contained in the given error.
func (err *ResponseError) StatusCode() int {
	return err.statusCode
}

// checkResponseErrorStatusCode checks that the given error is a ResponseError
// with given values.
// statusCode: This is always checked
// errorCode: This is checked if the given value is non-zero
// reason: This is checked if the given value is not empty
func checkResponseErrorStatusCode(err error, statusCode, errorCode int, reason string) bool {
	respErr, ok := err.(*ResponseError)
	if !ok {
		return false
	}

	if respErr.StatusCode() != statusCode {
		return false
	}

	if errorCode != 0 {
		// Check for a matching error code
		if respErr.ErrorCode() != errorCode {
			return false
		}
	}

	if reason != "" {
		// Check reason part of the status text.
		// See also util.go#newReason
		if (respErr.statusText != reason) && !strings.HasSuffix(respErr.statusText, ": "+reason) {
			return false
		}
	}

	return true
}

func IsResourceUpError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_UP, 0, "")
}

func IsResourceUpWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_UP, errorCode, "")
}

func IsResourceDownError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_DOWN, 0, "")
}

func IsResourceDownWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_DOWN, errorCode, "")
}

func IsResourceNotFoundError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_NOT_FOUND, 0, "")
}

func IsResourceNotFoundWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_NOT_FOUND, errorCode, "")
}

func IsResourceAlreadyExistsError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_ALREADY_EXISTS, 0, "")
}

func IsResourceAlreadyExistsWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_ALREADY_EXISTS, errorCode, "")
}

func IsResourceInvalidCredentialsError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_INVALID_CREDENTIALS, 0, "")
}

func IsResourceInvalidCredentialsWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_INVALID_CREDENTIALS, errorCode, "")
}

func IsWrongInputError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_WRONG_INPUT, 0, "")
}

func IsWrongInputWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_WRONG_INPUT, errorCode, "")
}

func IsWrongInputWithReasonError(err error, reason string) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_WRONG_INPUT, 0, reason)
}

func IsUserError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_USER_ERROR, 0, "")
}

func IsUserWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_USER_ERROR, errorCode, "")
}

func IsServerError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_SERVER_ERROR, 0, "")
}
func IsServerWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_SERVER_ERROR, errorCode, "")
}

func IsServerErrorWithReason(err error, reason string) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_SERVER_ERROR, 0, reason)
}

func IsInvalidVersionError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_INVALID_VERSION_ERROR, 0, "")
}

func IsInvalidVersionWithCode(err error, errorCode int) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_INVALID_VERSION_ERROR, errorCode, "")
}
