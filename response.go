package apischema

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/juju/errgo"
)

// The Response structure is a representation of responses
// that are send back in response to HTTP requests.
type Response struct {
	StatusCode int             `json:"status_code"`
	ErrorCode  int             `json:"error_code,omitempty"`
	StatusText string          `json:"status_text"`
	Data       json.RawMessage `json:"data,omitempty"`
}

// NewEmptyResponse creates a Response object with given values
func NewEmptyResponse(statusCode int, statusText string) *Response {
	return &Response{
		StatusCode: statusCode,
		StatusText: statusText,
	}
}

// SetStatusCode overwrites the StatusCode of the given response
// and returns the given Response object.
func (r *Response) SetStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	return r
}

// SetErrorCode overwrites the ErrorCode of the given response
// and returns the given Response object.
func (r *Response) SetErrorCode(errorCode int) *Response {
	r.ErrorCode = errorCode
	return r
}

// SetStatusText overwrites the StatusText of the given response
// and returns the given Response object.
func (r *Response) SetStatusText(statusText string) *Response {
	r.StatusText = statusText
	return r
}

// NewResponse creates a Response object with given values
func NewResponse(statusCode int, statusText string, data interface{}) (*Response, error) {
	rawData, err := json.Marshal(data)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return &Response{
		StatusCode: statusCode,
		StatusText: statusText,
		Data:       rawData,
	}, nil
}

// ParseResponse parses the given content into a Response object.
func ParseResponse(resBody *io.ReadCloser) (*Response, error) {
	byteSlice, err := ioutil.ReadAll(*resBody)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	// This is a hack to be able to read from response body twice. Because we
	// need to read the response body to identify the actual status of the
	// response, we consume the stream maybe somebody else would like too. Thus
	// we buffer the response body and write it back to the original response
	// body reference.
	*resBody = ioutil.NopCloser(bytes.NewReader(byteSlice))

	target := Response{}

	if isJSON(string(byteSlice)) {
		if err := json.Unmarshal(byteSlice, &target); err != nil {
			// In case we receive a response we did not expect and cannot read, we just
			// return an error containing the content of the response.
			return nil, newUnexpectedContentError(string(byteSlice))
		}
	}

	return &target, nil
}

// FromHTTPResponse parses the content in the body of the given HTTP response
// into a Response object.
func FromHTTPResponse(resp *http.Response, err error) (*Response, error) {
	if err != nil {
		return nil, err
	}
	return ParseResponse(&resp.Body)
}

// EnsureStatusCodes checks that the status code of the given response
// equals to one of the given status codes.
// If there is a match, nil is returned.
// If there is no match, a ResponseError is returned.
func (resp *Response) EnsureStatusCodes(statusCodes ...int) error {
	if containsInt(statusCodes, resp.StatusCode) {
		return nil
	}
	return NewResponseError(resp)
}

// UnmarshalData unmarshals the data field of the given Response
// into the given interface.
func (resp *Response) UnmarshalData(v interface{}) error {
	if err := json.Unmarshal(resp.Data, v); err != nil {
		// In case we receive a data field we did not expect, we just
		// return an error containing the content of the response.
		return newUnexpectedContentError(string(resp.Data))
	}

	return nil
}

func containsInt(list []int, n int) bool {
	for _, listN := range list {
		if listN == n {
			return true
		}
	}

	return false
}
