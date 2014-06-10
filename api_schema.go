package apischema

const (
  STATUS_CODE_DATA = iota
  STATUS_CODE_RESSOURCE_CREATED
  STATUS_CODE_RESSOURCE_ALREADY_EXISTS
)

type ResponsePayload struct {
	StatusCode int         `json:"status_code"`
	StatusText string      `json:"status_text"`
	Data       interface{} `json:"data,omitempty"`
}

// Responder method.
func StatusData(data interface{}) ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_DATA,
		StatusText: "success",
    Data: data,
	}
}

// Receiver method.
func IsStatusData(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_DATA, responseBody)
}

// Responder method.
func StatusRessourceCreated() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_CREATED,
		StatusText: "resource created",
	}
}

// Receiver method.
func IsStatusRessourceCreated(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_RESSOURCE_CREATED, responseBody)
}

// Responder method.
func StatusRessourceAlreadyExists() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_ALREADY_EXISTS,
		StatusText: "resource already exists",
	}
}

// Receiver method.
func IsStatusRessourceAlreadyExists(responseBody string) (bool, error) {
  return IsStatus(STATUS_CODE_RESSOURCE_ALREADY_EXISTS, responseBody)
}
