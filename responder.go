package apischema

func StatusData(data interface{}) ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_DATA,
		StatusText: "success",
    Data: data,
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

func StatusRessourceAlreadyExists() ResponsePayload {
	return ResponsePayload{
		StatusCode: STATUS_CODE_RESSOURCE_ALREADY_EXISTS,
		StatusText: "resource already exists",
	}
}
