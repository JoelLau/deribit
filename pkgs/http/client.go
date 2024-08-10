package http

import (
	"fmt"
	"net/http"
)

// Expects net/http's Client
type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type OptFunc func(Request) (Request, error)
type Request *http.Request
type Response *http.Response

// Thin wrapper to make http.Client more ergonomic
type HttpClient struct {
	Http HttpDoer
}

func NewHttpClient() (c *HttpClient) {
	c = &HttpClient{}
	c.Http = &http.Client{}

	return
}

func (client *HttpClient) Get(url string, opts ...OptFunc) (res Response, err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("error intialising http request: %w", err)
		return
	}

	for _, opt := range opts {
		if request, err = opt(request); err != nil {
			err = fmt.Errorf("error applying optfunc: %w", err)
			return
		}
	}

	if res, err = client.Http.Do(request); err != nil {
		err = fmt.Errorf("error performing http request: %w", err)
		return
	}

	return
}

func (client *HttpClient) Post(url string, body string, opts ...OptFunc) (res Response, err error) {
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return
	}

	for _, opt := range opts {
		if request, err = opt(request); err != nil {
			return
		}
	}

	if res, err = client.Http.Do(request); err != nil {
		return
	}

	return
}
