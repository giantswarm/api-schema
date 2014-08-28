package apischema

import (
	"io"
)

func IsStatusData(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_DATA, responseBody)
}

func IsStatusResourceUp(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_UP, responseBody)
}

func IsStatusResourceDown(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_DOWN, responseBody)
}

func IsStatusResourceCreated(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_CREATED, responseBody)
}

func IsStatusResourceStarted(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_STARTED, responseBody)
}

func IsStatusResourceStopped(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_STOPPED, responseBody)
}

func IsStatusResourceUpdated(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_UPDATED, responseBody)
}

func IsStatusResourceDeleted(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_DELETED, responseBody)
}

func IsStatusResourceNotFound(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_NOT_FOUND, responseBody)
}

func IsStatusResourceAlreadyExists(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_ALREADY_EXISTS, responseBody)
}

func IsStatusResourceInvalidCredentials(responseBody *io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESOURCE_INVALID_CREDENTIALS, responseBody)
}
