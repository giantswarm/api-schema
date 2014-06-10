package apischema

import (
  "net/http"
  "encoding/json"

  "github.com/juju/errgo"
)

func IsStatus(statusCode int, responseBody string) (bool, error) {
  var responsePayload ResponsePayload
  if err := json.Unmarshal([]byte(responseBody), &responsePayload); err != nil {
    return false, errgo.Mask(err)
  }

  if responsePayload.StatusCode == STATUS_CODE_DATA {
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

