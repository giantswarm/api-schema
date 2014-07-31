package apischema

import (
	"io"
)

func IsStatusData(responseBody io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_DATA, responseBody)
}

func IsStatusRessourceCreated(responseBody io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESSOURCE_CREATED, responseBody)
}

func IsStatusRessourceStarted(responseBody io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESSOURCE_STARTED, responseBody)
}

func IsStatusRessourceStopped(responseBody io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESSOURCE_STOPPED, responseBody)
}

func IsStatusRessourceUpdated(responseBody io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESSOURCE_UPDATED, responseBody)
}

func IsStatusRessourceDeleted(responseBody io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESSOURCE_DELETED, responseBody)
}

func IsStatusRessourceNotFound(responseBody io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESSOURCE_NOT_FOUND, responseBody)
}

func IsStatusRessourceAlreadyExists(responseBody io.ReadCloser) (bool, error) {
	return IsStatusWithRawBody(STATUS_CODE_RESSOURCE_ALREADY_EXISTS, responseBody)
}
