package apischema

type ResponseError struct {
	statusCode int
	statusText string
}

func NewResponseError(resp *Response) error {
	return &ResponseError{
		statusCode: resp.StatusCode,
		statusText: resp.StatusText,
	}
}

func (err *ResponseError) Error() string {
	return err.statusText
}

func (err *ResponseError) StatusCode() int {
	return err.statusCode
}

func checkResponseErrorStatusCode(err error, statusCode int) bool {
	respErr, ok := err.(*ResponseError)
	if !ok {
		return false
	}

	return respErr.StatusCode() == statusCode
}

func IsResourceUpError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_UP)
}

func IsResourceDownError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_DOWN)
}

func IsResourceNotFoundError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_NOT_FOUND)
}

func IsResourceAlreadyExistsError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_ALREADY_EXISTS)
}

func IsResourceInvalidCredentialsError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_RESOURCE_INVALID_CREDENTIALS)
}

func IsWrongInputError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_WRONG_INPUT)
}

func IsUserError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_USER_ERROR)
}

func IsServerError(err error) bool {
	return checkResponseErrorStatusCode(err, STATUS_CODE_SERVER_ERROR)
}
