package apischema

func StatusData(data interface{}) ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_DATA,
		StatusText: "success",
		Data:       data,
	}
}

func StatusResourceUp() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_UP,
		StatusText: "resource up",
	}
}

func StatusResourceDown() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_DOWN,
		StatusText: "resource down",
	}
}

func StatusResourceCreated() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_CREATED,
		StatusText: "resource created",
	}
}

func StatusResourceStarted() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_STARTED,
		StatusText: "resource started",
	}
}

func StatusResourceStopped() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_STOPPED,
		StatusText: "resource stopped",
	}
}

func StatusResourceUpdated() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_UPDATED,
		StatusText: "resource updated",
	}
}

func StatusResourceDeleted() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_DELETED,
		StatusText: "resource deleted",
	}
}

func StatusResourceNotFound() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_NOT_FOUND,
		StatusText: "resource not found",
	}
}

func StatusResourceAlreadyExists() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_ALREADY_EXISTS,
		StatusText: "resource already exists",
	}
}

func StatusResourceInvalidCredentials() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESOURCE_INVALID_CREDENTIALS,
		StatusText: "resource invalid credentials",
	}
}
