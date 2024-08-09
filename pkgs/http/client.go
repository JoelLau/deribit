package http

import (
	"net/http"
)

type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
type OptFunc func(*http.Request) (*http.Request, error)

// Thin wrapper to make http.Client more ergonomic
type HttpClient struct {
	http HttpDoer
}

func (client *HttpClient) Get(url string, opts ...OptFunc) (res *http.Response, err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	for _, opt := range opts {
		if request, err = opt(request); err != nil {
			return
		}
	}

	if res, err = client.http.Do(request); err != nil {
		return
	}

	return
}
