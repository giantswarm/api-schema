package apischema

func StatusData(data interface{}) ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_DATA,
		StatusText: "success",
		Data:       data,
	}
}

func StatusRessourceUp() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_UP,
		StatusText: "resource up",
	}
}

func StatusRessourceDown() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_DOWN,
		StatusText: "resource down",
	}
}

func StatusRessourceCreated() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_CREATED,
		StatusText: "resource created",
	}
}

func StatusRessourceStarted() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_STARTED,
		StatusText: "resource started",
	}
}

func StatusRessourceStopped() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_STOPPED,
		StatusText: "resource stopped",
	}
}

func StatusRessourceUpdated() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_UPDATED,
		StatusText: "resource updated",
	}
}

func StatusRessourceDeleted() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_DELETED,
		StatusText: "resource deleted",
	}
}

func StatusRessourceNotFound() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_NOT_FOUND,
		StatusText: "resource not found",
	}
}

func StatusRessourceAlreadyExists() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_ALREADY_EXISTS,
		StatusText: "resource already exists",
	}
}
