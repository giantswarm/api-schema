package apischema

const (
	STATUS_CODE_DATA = 10000

	STATUS_CODE_RESOURCE_UP   = 10001
	STATUS_CODE_RESOURCE_DOWN = 10002

	STATUS_CODE_RESOURCE_CREATED = 10003
	STATUS_CODE_RESOURCE_STARTED = 10004
	STATUS_CODE_RESOURCE_STOPPED = 10005
	STATUS_CODE_RESOURCE_UPDATED = 10006
	STATUS_CODE_RESOURCE_DELETED = 10007

	STATUS_CODE_RESOURCE_NOT_FOUND           = 10008
	STATUS_CODE_RESOURCE_ALREADY_EXISTS      = 10009
	STATUS_CODE_RESOURCE_INVALID_CREDENTIALS = 10010

	STATUS_CODE_CONDITION_TRUE  = 10011
	STATUS_CODE_CONDITION_FALSE = 10012
	STATUS_CODE_WRONG_INPUT     = 10013

	// NOTE: We want explicit values instead of using iota to be more robust
	// against reordering and/or adding new values inbetween.
)

type ResponsePayload struct {
	StatusCode int         `json:"status_code"`
	StatusText string      `json:"status_text"`
	Data       interface{} `json:"data,omitempty"`
}
