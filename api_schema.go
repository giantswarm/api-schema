package apischema

const (
	STATUS_CODE_DATA = 10000 + iota

	STATUS_CODE_RESOURCE_UP
	STATUS_CODE_RESOURCE_DOWN

	STATUS_CODE_RESOURCE_CREATED
	STATUS_CODE_RESOURCE_STARTED
	STATUS_CODE_RESOURCE_STOPPED
	STATUS_CODE_RESOURCE_UPDATED
	STATUS_CODE_RESOURCE_DELETED

	STATUS_CODE_RESOURCE_NOT_FOUND
	STATUS_CODE_RESOURCE_ALREADY_EXISTS
)

type ResponsePayload struct {
	StatusCode int         `json:"status_code"`
	StatusText string      `json:"status_text"`
	Data       interface{} `json:"data,omitempty"`
}
