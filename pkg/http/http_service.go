package http

import (
	"errors"
	"io"
	"net/http"

	"github.com/tanveerprottoy/rest-api-starter-go/pkg/adapter"
)

func Request[T any](
	method string,
	url string,
	header http.Header,
	body io.Reader,
	httpClient *HTTPClient,
) (*T, error) {
	code, resBody, err := httpClient.Request(method, url, header, body)
	if err != nil {
		return nil, err
	}
	if code >= http.StatusOK && code < http.StatusMultipleChoices {
		// res ok, parse response body to type
		d, err := adapter.BytesToValue[T](resBody)
		if err != nil {
			return nil, err
		}
		return d, nil
	} else {
		// res not ok, parse error
		errBody, err := adapter.BytesToValue[ErrorBody](resBody)
		if err != nil {
			return nil, err
		}
		errMessage := "Something went wrong"
		if errBody.Message != "" {
			errMessage = errBody.Message
		} else if errBody.Error != "" {
			errMessage = errBody.Error
		}
		return nil, errors.New(errMessage)
	}
}
