package httppkg

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient struct {
	Client *http.Client
}

func NewHTTPClient(
	timeout time.Duration,
	transport *http.Transport,
	checkRedirectFunc func(req *http.Request, via []*http.Request) error,
) *HTTPClient {
	c := new(HTTPClient)
	c.Client = &http.Client{
		Timeout: timeout,
	}
	if transport != nil {
		c.Client.Transport = transport
	}
	if checkRedirectFunc != nil {
		c.Client.CheckRedirect = checkRedirectFunc
	}
	return c
}

func (c *HTTPClient) Request(
	method string,
	url string,
	header http.Header,
	body io.Reader,
) (int, []byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return -1, nil, err
	}
	if header != nil {
		req.Header = header
	}
	res, err := c.Client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return -1, nil, err
	}
	return res.StatusCode, resBody, nil
}

// ex:
// resp, err := http.PostForm("http://example.com/form",
// url.Values{"key": {"Value"}, "id": {"123"}})
func (c *HTTPClient) PostForm(
	url string,
	header http.Header,
	values url.Values,
) (int, []byte, error) {
	res, err := http.PostForm(url, values)
	if err != nil {
		return -1, nil, err
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return -1, nil, err
	}
	return res.StatusCode, resBody, nil
}
