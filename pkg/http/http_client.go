package http

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	Client *http.Client
}

func NewHTTPClient(
	timeout time.Duration,
	checkRedirectFunc func(req *http.Request, via []*http.Request) error,
	transport *http.Transport,
) *HTTPClient {
	c := new(HTTPClient)
	c.Client = &http.Client{
		Timeout:       timeout,
		CheckRedirect: checkRedirectFunc,
		Transport:     transport,
	}
	return c
}

func (c *HTTPClient) Request(
	method string,
	url string,
	header http.Header,
	body io.Reader,
) any {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil
	}
	if header != nil {
		req.Header = header
	}
	res, err := c.Client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil
	}
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil
	}
	fmt.Printf("client: response body: %s\n", resBody)
	return resBody
}
