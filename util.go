package apischema

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/juju/errgo"
)

func IsStatus(statusCode int, responseBody string) (bool, error) {
	return isStatus(statusCode, responseBody)
}

func IsStatusWithRawBody(statusCode int, responseBody io.ReadCloser) (bool, error) {
	bytes, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return false, errgo.Mask(err, errgo.Any)
	}

	return isStatus(statusCode, string(bytes))
}

func isStatus(statusCode int, responseBody string) (bool, error) {
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
