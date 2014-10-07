package apischema

func StatusData(data interface{}) ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_DATA,
		StatusText: "success",
		Data:       data,
	}
}

func StatusResourceUp() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_UP, "resource up")
}

func StatusResourceDown() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_DOWN, "resource down")
}

func StatusResourceCreated() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_CREATED, "resource created")
}

func StatusResourceStarted() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_STARTED, "resource started")
}

func StatusResourceStopped() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_STOPPED, "resource stopped")
}

func StatusResourceUpdated() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_UPDATED, "resource updated")
}

func StatusResourceDeleted() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_DELETED, "resource deleted")
}

func StatusResourceNotFound() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_NOT_FOUND, "resource not found")
}

func StatusResourceAlreadyExists() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_ALREADY_EXISTS, "resource already exists")
}

func StatusResourceInvalidCredentials() *Response {
	return NewEmptyResponse(STATUS_CODE_RESOURCE_INVALID_CREDENTIALS, "resource invalid credentials")
}

func StatusWrongInput() *Response {
	return StatusWrongInputWithReason("")
}

func StatusWrongInputWithReason(reason string) *Response {
	text := "wrong input"

	if len(reason) > 0 {
		text = text + ": " + reason
	}

	return NewEmptyResponse(STATUS_CODE_WRONG_INPUT, text)
}
