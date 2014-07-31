package apischema

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/juju/errgo"
)

func IsStatus(statusCode int, responseBody string) (bool, error) {
	return isStatus(statusCode, responseBody)
}

func IsStatusWithRawBody(statusCode int, responseBody *io.ReadCloser) (bool, error) {
	byteSlice, err := ioutil.ReadAll(*responseBody)
	if err != nil {
		return false, errgo.Mask(err, errgo.Any)
	}

	// This is a hack to be able to read from response body twice. Because we
	// need to read the response body to identify the actual status of the
	// response, we consume the stream maybe somebody else would like too. Thus
	// we buffer the response body and write it back to the original response
	// body reference.
	*responseBody = ioutil.NopCloser(bytes.NewReader(byteSlice))

	return isStatus(statusCode, string(byteSlice))
}

func isStatus(statusCode int, responseBody string) (bool, error) {
	var responsePayload ResponsePayload
	if err := json.Unmarshal([]byte(responseBody), &responsePayload); err != nil {
		return false, errgo.Newf("Invalid request body. Expecting json. Aborting...")
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
