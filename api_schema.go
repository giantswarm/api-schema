package apischema

const (
  STATUS_CODE_DATA = 10000 + iota
  STATUS_CODE_RESSOURCE_CREATED
  STATUS_CODE_RESSOURCE_STARTED
  STATUS_CODE_RESSOURCE_STOPPED
  STATUS_CODE_RESSOURCE_UPDATED
  STATUS_CODE_RESSOURCE_DELETED
  STATUS_CODE_RESSOURCE_ALREADY_EXISTS
)

type ResponsePayload struct {
	StatusCode int         `json:"status_code"`
	StatusText string      `json:"status_text"`
	Data       interface{} `json:"data,omitempty"`
}
