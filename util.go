package apischema

import (
	"encoding/json"
	"net/http"

	"github.com/juju/errgo"
)

func IsStatus(statusCode int, responseBody string) (bool, error) {
	var responsePayload ResponsePayload
	if err := json.Unmarshal([]byte(responseBody), &responsePayload); err != nil {
		return false, errgo.Mask(err)
	}

	if responsePayload.StatusCode == statusCode {
		return true, nil
	}

	return false, nil
}

func IsSuccessResponse(statusCode int) bool {
	return statusCode == http.StatusOK
}

func IsFailureResponse(statusCode int) bool {
	return statusCode == http.StatusInternalServerError
}
